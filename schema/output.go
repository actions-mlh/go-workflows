package generate

import (
	// "bytes"
	"fmt"
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

func getOrderedStructNames(m map[string]Struct) []string {
	keys := make([]string, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	return keys
}

// Output generates code and writes to w.
func Output(w io.Writer, g *Generator, pkg string) {
	structs := g.Structs

	fmt.Fprintln(w, "// Code generated by schema-generate. DO NOT EDIT.")
	fmt.Fprintf(w, "package %v\n", pkg)

	// write all the code into a buffer, compiler functions will return list of imports
	// write list of imports into main output stream, followed by the code

	imports := []string{
		"fmt",
		"gopkg.in/yaml.v3",
	}

	if len(imports) > 0 {
		fmt.Fprint(w, "import (")
		for _, importfile := range imports {
			fmt.Fprintf(w, "\n\t\"%s\"", importfile)
		}
		fmt.Fprintln(w, "\n)")
	}

	for _, k := range getOrderedStructNames(structs) {
		s := structs[k]
		// codeBuf := new(bytes.Buffer)
		fmt.Fprintln(w, "")
		outputNameAndDescriptionComment(s.Name, s.Description, w)
		fmt.Fprintf(w, "type %s struct {\n", s.Name)

		for _, f := range getOrderedFields(s.Fields) {
			// Only apply omitempty if the field is not required.
			omitempty := ",omitempty"
			if f.Required {
				omitempty = ""
			}

			if f.Description != "" {
				outputFieldDescriptionComment(f.Description, w)
			}
			fmt.Fprintf(w, "\t%s *%s `yaml:\"%s%s\"`\n", f.Name, f.Type, f.YAMLName, omitempty)
		}
		fmt.Fprintf(w, "\tRaw *yaml.Node\n")
		fmt.Fprintln(w, "}")
		emitUnMarshalCode(w, s)
	}
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

func outputNameAndDescriptionComment(name, description string, w io.Writer) {
	if strings.Index(description, "\n") == -1 {
		fmt.Fprintf(w, "// %s %s\n", name, description)
		return
	}

	dl := strings.Split(description, "\n")
	fmt.Fprintf(w, "// %s %s\n", name, strings.Join(dl, "\n// "))
}

func outputFieldDescriptionComment(description string, w io.Writer) {
	if strings.Index(description, "\n") == -1 {
		fmt.Fprintf(w, "\n  // %s\n", description)
		return
	}

	dl := strings.Split(description, "\n")
	fmt.Fprintf(w, "\n  // %s\n", strings.Join(dl, "\n  // "))
}

func cleanPackageName(pkg string) string {
	pkg = strings.Replace(pkg, ".", "", -1)
	pkg = strings.Replace(pkg, "_", "", -1)
	pkg = strings.Replace(pkg, "-", "", -1)
	return pkg
}
