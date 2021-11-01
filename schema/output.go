package generate

import (
	"io"
	"sort"
	"strings"
	"text/template"
)

func getOrderedFields(m map[string]Field) []Field {
	fields := make([]Field, len(m))
	idx := 0
	for _, v := range m {
		fields[idx] = v
		idx++
	}
	sort.Slice(fields, func (i, j int) bool {
			return fields[i].Type < fields[j].Type
		})
	return fields
}

func getOrderedStructs(m map[string]Struct) []Struct {
	structs := make([]Struct, len(m))
	idx := 0
	for _, v := range m {
		structs[idx] = v
		idx++
	}
	sort.Slice(structs, func (i, j int) bool {
			return structs[i].Name < structs[j].Name
		})
	return structs
}

// Output generates code and writes to w.
func Output(w io.Writer, g *Generator, pkg string) error {
	templateText := `// Code generated by schema-generate. DO NOT EDIT.
package {{pkg}}
import (
	"fmt"
	"gopkg.in/yaml.v3"
)
{{range . | getOrderedStructs}}
// {{.Name}} {{Replace .Description "\n" "\n// " -1}}
type {{.Name}} struct {
{{range .Fields | getOrderedFields}}{{if .Description}}
  // {{Replace .Description "\n" "\n  // " -1}}
{{end}}	{{.Name}} *{{.Type}} `+"`"+`yaml:"{{.YAMLName}}{{if not .Required}},omitempty{{end}}"`+"`"+`
{{end}}	Raw *yaml.Node
}

func (node *{{.Name}}) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for i := 0; i < len(value.Content); i++ {
		nodeName := value.Content[i]
		switch nodeName.Value {
			{{range .Fields | getOrderedFields}}
			case "{{ToLower .Name}}":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.{{.Name}} = new({{.Type}})
				err := nodeValue.Decode(node.{{.Name}})
				if err != nil {
					return err
				}
			{{end}}
		}
	}
	return nil
}
{{end}}
`
	funcMap := template.FuncMap{
		"getOrderedStructs": getOrderedStructs,
		"getOrderedFields": getOrderedFields,
		"ToLower": strings.ToLower,
		"Replace": strings.Replace,
		"pkg": func () string { return pkg },
	}
	t, err := template.New("gen_schema.go").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err
	}
	return t.Execute(w, g.Structs)
}

func emitUnMarshalCode(w io.Writer, s Struct) error {
	templateText := `
func (node *{{.Name}}) UnmarshalYAML(value *yaml.Node) error {
	node.Raw = value
	for i := 0; i < len(value.Content); i++ {
		nodeName := value.Content[i]
		switch nodeName.Value {
			{{range .Fields | getOrderedFields}}
			case "{{ToLower .Name}}":
				i++
				if i >= len(value.Content) {
					return fmt.Errorf("value.Content mismatch")
				}
				nodeValue := value.Content[i]
				node.{{.Name}} = new({{.Type}})
				err := nodeValue.Decode(node.{{.Name}})
				if err != nil {
					return err
				}
			{{end}}
		}
	}
	return nil
}
`
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
		"getOrderedFields": getOrderedFields,
	}
	t, err := template.New("unmarshal").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err
	}
	return t.Execute(w, s)
}
