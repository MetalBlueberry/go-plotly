package generator

import (
	"embed"
	"fmt"
	"io"
	"sort"
	"text/template"

	"github.com/huandu/xstrings"
)

type Creator interface {
	Create(name string) (io.WriteCloser, error)
}

type Renderer struct {
	tmpl *template.Template
	root *Root

	fs Creator
}

//go:embed templates/*.tmpl
var templates embed.FS

func NewRenderer(fs Creator, root *Root) (*Renderer, error) {
	r := &Renderer{
		root: root,
		fs:   fs,
	}
	tmpl, err := template.New("base").ParseFS(templates, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}
	r.tmpl = tmpl
	return r, nil
}

func (r *Renderer) CreateTrace(traceName string) error {
	traceFile, err := r.fs.Create(traceName + ".go")
	if err != nil {
		return err
	}
	defer traceFile.Close()
	r.WriteTrace(traceName, traceFile)

	return nil
}

func sortKeys(attr map[string]*Attribute) []string {
	keys := make([]string, 0, len(attr))
	for k := range attr {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (r *Renderer) parseAttributes(prefix string, attr map[string]*Attribute) ([]StructField, []Struct, error) {
	fields := []StructField{}
	structs := []Struct{}

	for _, name := range sortKeys(attr) {
		if name == "_deprecated" {
			continue
		}

		attr := attr[name]

		switch {
		case attr.Role == "object":
			obj, err := r.parseObject(prefix, attr)
			if err != nil {
				return nil, nil, fmt.Errorf("cannot parse object %s, %w", attr.Name, err)
			}
			structs = append(structs, obj...)
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        obj[0].Name,
				Description: attr.Description,
			})

		case attr.ValType == "flaglist":
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        prefix + xstrings.ToCamelCase(attr.Name),
				Description: attr.Description,
			})

		default:
			ty := ValTypeMap[attr.ValType]
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        ty,
				Description: attr.Description + " " + attr.ValType,
			})
		}
	}
	return fields, structs, nil
}

func (r *Renderer) parseObject(prefix string, attr *Attribute) ([]Struct, error) {
	structs := []Struct{}

	objStruct := Struct{
		Name:        prefix + xstrings.ToCamelCase(attr.Name),
		Description: attr.Description,
		Fields:      []StructField{},
	}

	fields, structs, err := r.parseAttributes(objStruct.Name, attr.Attributes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse attributes, %w", err)
	}
	objStruct.Fields = fields

	structs = append([]Struct{objStruct}, structs...)
	return structs, nil
}

func (r *Renderer) WriteTrace(traceName string, w io.Writer) error {
	fmt.Fprint(w, "package grob")

	trace := r.root.Schema.Traces[traceName]
	traceStruct := Struct{
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
	fields, objectStructs, err := r.parseAttributes(traceStruct.Name, trace.Attributes.Names)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	traceStruct.Fields = append(traceStruct.Fields, fields...)

	err = r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceStruct)
	if err != nil {
		return err
	}
	for i := range objectStructs {
		err := r.tmpl.ExecuteTemplate(w, "trace.tmpl", objectStructs[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Renderer) WriteTraces(w io.Writer) error {
	panic("Not Implemented Yet")
	fmt.Fprintln(w, "package grob")
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

var ValTypeMap = map[string]string{
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
