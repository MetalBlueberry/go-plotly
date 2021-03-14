package generator

import (
	"fmt"
	"io"
	"text/template"

	"github.com/huandu/xstrings"
)

type Renderer struct {
	tmpl *template.Template
	root *Root
}

func NewRenderer(glob string, root *Root) (*Renderer, error) {
	r := &Renderer{
		root: root,
	}
	tmpl, err := template.New("base").
		ParseGlob(glob)
	if err != nil {
		return nil, err
	}
	r.tmpl = tmpl
	return r, nil
}

func (r *Renderer) WriteTraces(w io.Writer) error {
	fmt.Fprintln(w, "package grob")

	for _, key := range r.root.Schema.Traces.Sorted() {
		trace := r.root.Schema.Traces[key]
		s := Struct{
			Name:        xstrings.ToCamelCase(trace.Type),
			Description: trace.Meta.Description,
			Fields: []StructField{
				{
					Name:        "Type",
					JSONName:    "type",
					Type:        "TraceType",
					Description: "is the type of the plot",
				},
			},
		}
		for _, name := range trace.Attributes.Sorted() {
			if name == "_deprecated" {
				continue
			}
			attr := trace.Attributes.Names[name]
			if attr.ValType == "flaglist" {
				s.Fields = append(s.Fields, StructField{
					Name:        xstrings.ToCamelCase(attr.Name),
					JSONName:    attr.Name,
					Type:        xstrings.ToCamelCase(trace.Type + "_" + attr.Name),
					Description: attr.Description,
				})
			} else {
				ty := TypeMap[attr.ValType]
				s.Fields = append(s.Fields, StructField{
					Name:        xstrings.ToCamelCase(attr.Name),
					JSONName:    attr.Name,
					Type:        ty,
					Description: attr.Description,
				})
			}
		}
		err := r.tmpl.ExecuteTemplate(w, "trace.tmpl", s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Renderer) WriteFlaglist(w io.Writer) error {
	fmt.Fprintln(w, "package grob")

	return nil
}

type Struct struct {
	Name        string
	Description string
	Fields      []StructField
}
type StructField struct {
	Name        string
	JSONName    string
	Type        string
	Description string
}

var TypeMap = map[string]string{
	"info_array": "interface{}",
	"subplotid":  "String",
	"data_array": "interface{}",
	"any":        "interface{}",
	"angle":      "float64",
	"color":      "String",
	"number":     "float64",
	"string":     "String",
	"integer":    "int64",
	"boolean":    "Bool",
}
