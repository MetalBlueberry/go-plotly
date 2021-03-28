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

func (r *TraceFile) parseAttributes(prefix string, attr map[string]*Attribute) ([]StructField, error) {
	fields := make([]StructField, 0, len(attr))

	for _, name := range sortKeys(attr) {
		if name == "_deprecated" {
			continue
		}

		attr := attr[name]

		switch {
		case attr.Role == "object":
			name := prefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseObject(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse object %s, %w", attr.Name, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        name,
				Description: attr.ValType + " " + attr.Description,
			})

		case attr.ValType == "flaglist":
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        prefix + xstrings.ToCamelCase(attr.Name),
				Description: attr.ValType + " " + attr.Description,
			})

		default:
			ty := ValTypeMap[attr.ValType]
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        ty,
				Description: attr.ValType + " " + attr.Description,
			})
		}
	}
	return fields, nil
}

func (r *TraceFile) parseObject(name string, attr *Attribute) error {
	objStruct := Struct{
		Name:        name,
		Description: attr.Description,
		Fields:      []StructField{},
	}

	fields, err := r.parseAttributes(objStruct.Name, attr.Attributes)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	objStruct.Fields = fields

	r.Objects = append(r.Objects, objStruct)
	return nil
}

func (r *Renderer) WriteTrace(traceName string, w io.Writer) error {
	trace := r.root.Schema.Traces[traceName]

	traceFile := TraceFile{
		Trace: Struct{
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
		},
		Objects: []Struct{},
		Enums:   []Enum{},
	}

	fields, err := traceFile.parseAttributes(traceFile.Trace.Name, trace.Attributes.Names)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	traceFile.Trace.Fields = append(traceFile.Trace.Fields, fields...)

	fmt.Fprintf(w, `package grob

var TraceType%s TraceType = "%s"
`, traceFile.Trace.Name, traceName)

	err = r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceFile.Trace)
	if err != nil {
		return err
	}
	for i := range traceFile.Objects {
		err := r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceFile.Objects[i])
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

type TraceFile struct {
	Trace   Struct
	Objects []Struct
	Enums   []Enum
}

type Enum struct {
	Name        string
	Description string
	Values      []string
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
