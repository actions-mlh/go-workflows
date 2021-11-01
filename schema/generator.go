package generate

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Generator will produce structs from the JSON schema.
type Generator struct {
	schemas  []*Schema
	resolver *RefResolver
	Structs  map[string]Struct
	Aliases  map[string]Field
	// cache for reference types; k=url v=type
	refs      map[string]string
	anonCount int
}

// Struct defines the data required to generate a struct in Go.
type Struct struct {
	// The ID within the JSON schema, e.g. #/definitions/address
	ID string
	// The golang name, e.g. "Address"
	Name string
	// Description of the struct
	Description string
	Fields      map[string]Field

	GenerateCode   bool
	AdditionalType string
}

// Field defines the data required to generate a field in Go.
type Field struct {
	// The golang name, e.g. "Address1"
	Name string
	// The YAML name, e.g. "address1"
	YAMLName string
	// The golang type of the field, e.g. a built-in type like "string" or the name of a struct generated
	// from the JSON schema.
	Type string
	// Required is set to true when the field is required.
	Required    bool
	Description string
}

// New creates an instance of a generator which will produce structs.
func New(schemas ...*Schema) *Generator {
	return &Generator{
		schemas:  schemas,
		resolver: NewRefResolver(schemas),
		Structs:  make(map[string]Struct),
		Aliases:  make(map[string]Field),
		refs:     make(map[string]string),
	}
}

// CreateTypes creates types from the JSON schemas, keyed by the golang name.
func (g *Generator) CreateTypes() (err error) {
	if err := g.resolver.Init(); err != nil {
		return err
	}
	// extract the types
	for _, schema := range g.schemas {
		name := g.getSchemaName("", schema)
		rootType, err := g.processSchema(name, schema)
		if err != nil {
			return err
		}
		// ugh: if it was anything but a struct the type will not be the name...
		if rootType != "*"+name {
			a := Field{
				Name:        name,
				YAMLName:    "",
				Type:        rootType,
				Required:    false,
				Description: schema.Description,
			}
			g.Aliases[a.Name] = a
		}
	}
	return
}

// returns the type refered to by schema after resolving all dependencies
func (g *Generator) processSchema(name string, schema *Schema) (string, error) {
	if schema.Type == "" {
		if schema.Items != nil {
			schema.Type = "array"
		} else {
			schema.Type = "object"
		}
	}
	
	strct := Struct{
		ID:          schema.ID(),
		Name:        name,
		Description: schema.Description,
		Fields:      make(map[string]Field, len(schema.Properties)),
	}

	if schema.Type != "array" && schema.Type != "object" {
		// typeName, err := getPrimitiveTypeName(schema.Type, "", false)
		// if err != nil {
		// 	return "", err
		// }
		s := strings.Split(name, "_")
		yamlName := strings.ToLower(s[len(s)-1])
		f := Field{
			Name:        "Value",
			YAMLName:    yamlName,
			Type:        schema.Type,
			Required:    contains(schema.Required, "Value"),
			Description: schema.Description,
		}
		if f.Required {
			strct.GenerateCode = true
		}
		strct.Fields[f.Name] = f
	}
	// definitions
	for key, subSchema := range schema.Definitions {
		fieldName := getGolangName(key)
		// calculate sub-schema name here, may not actually be used depending on type of schema!
		subSchemaName := g.getSchemaName(fieldName, subSchema)
		fieldType, err := g.processSchema(name + "_Definitions_" + subSchemaName, subSchema)
		if err != nil {
			return "", err
		}
		f := Field{
			Name:        fieldName,
			YAMLName:    key,
			Type:        fieldType,
			Required:    contains(schema.Required, key),
			Description: subSchema.Description,
		}
		if f.Required {
			strct.GenerateCode = true
		}
		strct.Fields[f.Name] = f
	}
	// cache the object name in case any sub-schemas recursively reference it
	schema.GeneratedType = "*" + name
	for propKey, prop := range schema.Properties {
		fieldName := getGolangName(propKey)
		// calculate sub-schema name here, may not actually be used depending on type of schema!
		subSchemaName := g.getSchemaName(fieldName, prop)
		fieldType, err := g.processSchema(name + "_Properties_" + subSchemaName, prop)
		if err != nil {
			return "", err
		}
		f := Field{
			Name:        fieldName,
			YAMLName:    propKey,
			Type:        fieldType,
			Required:    contains(schema.Required, propKey),
			Description: prop.Description,
		}
		if f.Required {
			strct.GenerateCode = true
		}
		strct.Fields[f.Name] = f
	}
	if schema.Items != nil {
		// subType: fallback name in case this array contains inline object without a title
		subName := g.getSchemaName(name, (*Schema)(schema.Items))
		s := strings.Split(subName, "_")
		yamlName := strings.ToLower(s[len(s)-1])
		subTyp, err := g.processSchema(name + "_Items_" + subName, (*Schema)(schema.Items))
		if err != nil {
			return "", err
		}
		finalType, err := getPrimitiveTypeName("array", subTyp, true)
		if err != nil {
			return "", err
		}
		f := Field{
			Name:        subName,
			YAMLName:    yamlName,
			Type:        finalType,
			Required:    contains(schema.Required, subName),
			Description: schema.Description,
		}
		if f.Required {
			strct.GenerateCode = true
		}
		strct.Fields[f.Name] = f
	}

	if schema.Reference != "" {
		schemaPath := g.resolver.GetPath(schema)
		if schema.Reference == "" {
			return "", errors.New("processReference empty reference: " + schemaPath)
		}
		refSchema, err := g.resolver.GetSchemaByReference(schema)
		if err != nil {
			return "", errors.New("processReference: reference \"" + schema.Reference + "\" not found at \"" + schemaPath + "\"")
		}
		if refSchema.GeneratedType == "" {
			// reference is not resolved yet. Do that now.
			refSchemaName := g.getSchemaName("", refSchema)
			typeName, err := g.processSchema("Root_Definitions_" + refSchemaName, refSchema)
			if err != nil {
				return "", err
			}
			f := Field{
				Name: "Ref",
				YAMLName: strings.ToLower(refSchemaName), // change to something appropriate when you know what it is
				Type: typeName,
				Required: contains(schema.Required, refSchemaName),
				Description: refSchema.Description,
			}
			if f.Required {
				strct.GenerateCode = true
			}
			strct.Fields[f.Name] = f
		}
	}
	
	g.Structs[name] = strct
	return name, nil
}

// util --------------------------------

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getPrimitiveTypeName(schemaType string, subType string, pointer bool) (name string, err error) {
	switch schemaType {
	case "array":
		if subType == "" {
			return "error_creating_array", errors.New("can't create an array of an empty subtype")
		}
		return "[]" + subType, nil
	case "boolean":
		return "bool", nil
	case "integer":
		return "int", nil
	case "number":
		return "float64", nil
	case "null":
		return "nil", nil
	case "object":
		if subType == "" {
			return "error_creating_object", errors.New("can't create an object of an empty subtype")
		}
		if pointer {
			return "*" + subType, nil
		}
		return subType, nil
	case "string":
		return "string", nil
	}

	return "undefined", fmt.Errorf("failed to get a primitive type for schemaType %s and subtype %s",
		schemaType, subType)
}

// return a name for this (sub-)schema.
func (g *Generator) getSchemaName(keyName string, schema *Schema) string {
	if len(schema.Title) > 0 {
		return getGolangName(schema.Title)
	}
	if keyName != "" {
		return getGolangName(keyName)
	}
	if schema.Parent == nil {
		return "Root"
	}
	if schema.JSONKey != "" {
		return getGolangName(schema.JSONKey)
	}
	if schema.Parent != nil && schema.Parent.JSONKey != "" {
		return getGolangName(schema.Parent.JSONKey + "_" + "Item")
	}
	g.anonCount++
	return fmt.Sprintf("Anonymous_%d", g.anonCount)
}

// getGolangName strips invalid characters out of golang struct or field names.
func getGolangName(s string) string {
	buf := bytes.NewBuffer([]byte{})
	for i, v := range splitOnAll(s, isNotAGoNameCharacter) {
		if i == 0 && strings.IndexAny(v, "0123456789") == 0 {
			// Go types are not allowed to start with a number, lets prefix with an underscore.
			buf.WriteRune('_')
		}
		if i > 0 {
			buf.WriteRune('_')
		}
		buf.WriteString(capitaliseFirstLetter(v))
	}
	return buf.String()
}

func splitOnAll(s string, shouldSplit func(r rune) bool) []string {
	rv := []string{}
	buf := bytes.NewBuffer([]byte{})
	for _, c := range s {
		if shouldSplit(c) {
			rv = append(rv, buf.String())
			buf.Reset()
		} else {
			buf.WriteRune(c)
		}
	}
	if buf.Len() > 0 {
		rv = append(rv, buf.String())
	}
	return rv
}

func isNotAGoNameCharacter(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return false
	}
	return true
}

func capitaliseFirstLetter(s string) string {
	if s == "" {
		return s
	}
	prefix := s[0:1]
	suffix := s[1:]
	return strings.ToUpper(prefix) + suffix
}
