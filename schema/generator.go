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
	Type string
	// The original schema from which this struct was built
	Schema *Schema
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
		refs:     make(map[string]string),
	}
}

// CreateTypes creates types from the JSON schemas, keyed by the golang name.
func (g *Generator) CreateTypes() error {
	if err := g.resolver.Init(); err != nil {
		return err
	}
	// extract the types

	strct, err := g.processSchema("Root", g.schemas[0])
	if err != nil {
		return err
	}
	g.Structs[strct.Name] = *strct
	return nil
}

// returns the type refered to by schema after resolving all dependencies
func (g *Generator) processSchema(name string, schema *Schema) (*Struct, error) {
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
		Type: schema.Type,
		Schema: schema,
	}

	// cache the object name in case any sub-schemas recursively reference it
	schema.GeneratedType = "*" + name

	// handle additional properties
	if schema.AdditionalProperties != nil &&
		(schema.AdditionalProperties.AdditionalPropertiesBool == nil ||
			*schema.AdditionalProperties.AdditionalPropertiesBool){
		schema.Type = schema.AdditionalProperties.Type
		schema.Reference = schema.AdditionalProperties.Reference
		schema.OneOf = schema.AdditionalProperties.OneOf
		schema.AnyOf = schema.AdditionalProperties.AnyOf
		schema.AllOf = schema.AdditionalProperties.AllOf
	}

	// process properties
	if len(schema.Properties) > 0 {
		err := g.processProperties(schema, strct, schema.Properties)
		if err != nil {
			return nil, err
		}
	}

	// TODO: add patternProperties
	return &strct, nil
}

func (g *Generator) processProperties(schema *Schema, strct Struct, properties map[string]*Schema) error {
	for propKey, prop := range properties {
		f, err := g.processField(propKey, prop, strct.Name)
		if err != nil {
			return err
		}
		f.Required = contains(schema.Required, propKey)
		strct.Fields[f.Name] = *f
	}
	return nil
}

func (g *Generator) processField(fieldName string, fieldSchema *Schema, parentName string) (*Field, error) {
	f := Field{
		Name:        getGolangName(fieldName),
		YAMLName:    fieldName,
		Description: fieldSchema.Description,
	}
	var newStruct *Struct
	var err error
	if fieldSchema.Reference != "" {
		refSchema, err := g.resolver.GetSchemaByReference(fieldSchema)
		if err != nil {
			return nil, errors.New("processField: reference \"" + fieldSchema.Reference + "\" not found at \"" + g.resolver.GetPath(fieldSchema) + "\"")
		}
		fieldSchema = refSchema
	}
	newStruct, err = g.processSchema(parentName + "_" + f.Name, fieldSchema)
	if err != nil {
		return nil, err
	}
	if newStruct.Type != "array" && newStruct.Type != "object" {
		f.Type = newStruct.Type
	} else if fieldSchema.OneOf != nil || fieldSchema.AllOf != nil || fieldSchema.AnyOf != nil {
		f.Type = "interface{}"
	} else {
		f.Type = newStruct.Name
		g.Structs[newStruct.Name] = *newStruct
	}
	return &f, nil
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
/*
func (g *Generator) processDefinitions(definitions map[string]*Schema) error {
	defStruct := Struct{
		ID: "#/definitions",
		Name: "Definitions",
		Description: "",
		Fields: make(map[string]Field, len(definitions)),
	}
	for key, schema := range definitions {
		fieldName := getGolangName(key)
		// calculate sub-schema name here, may not actually be used depending on type of schema!
		subSchemaName := g.getSchemaName(fieldName, schema)
		strct, err := g.processSchema("Definitions_" + subSchemaName, schema)
		if err != nil {
			return err
		}
		g.Structs[strct.Name] = *strct
		f := Field{
			Name:        fieldName,
			YAMLName:    key,
			Type:        strct.Name,
			Required:    contains(schema.Required, key),
			Description: schema.Description,
		}
		if f.Required {
			defStruct.GenerateCode = true
		}
		defStruct.Fields[f.Name] = f
	}
	g.Structs["Definitions"] = defStruct
	return nil
}
*/
