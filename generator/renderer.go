package generator

import (
	"embed"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
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
	err = r.WriteTrace(traceName, traceFile)
	if err != nil {
		return err
	}

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
		case attr.Role == RoleObject && len(attr.Items) > 0:
			fields = append(fields, StructField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     "interface{}",
				Description: []string{
					"It's an items array and what goes inside it's... messy... check the docs",
					"I will be happy if you want to contribute by implementing this",
					"just raise an issue before you start so we do not overlap",
				},
			})

		case attr.Role == RoleObject:
			name := prefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseObject(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse object %s, %w", name, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        name,
				Description: []string{string(attr.ValType) + " " + attr.Description},
			})

		case attr.ValType == ValTypeFlagList:
			name := prefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseFlaglist(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse flaglist %s, %w", name, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        prefix + xstrings.ToCamelCase(attr.Name),
				Description: []string{string(attr.ValType) + " " + attr.Description},
			})

		case attr.ValType == ValTypeEnum:
			name := prefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseEnum(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse enum %s, %w", name, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        name,
				Description: []string{string(attr.ValType) + " " + attr.Description},
			})

		// case attr.ValType == ValTypeColor
		default:
			ty := ValTypeMap[attr.ValType]
			fields = append(fields, StructField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     ty,
				Description: []string{
					string(attr.ValType),
					attr.Description,
				},
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

func (r *TraceFile) parseEnum(name string, attr *Attribute) error {

	values := make([]EnumValue, 0, len(attr.Values))
	for _, attrValue := range attr.Values {
		switch v := attrValue.(type) {
		case string:
			if v == "" {
				values = append(values, EnumValue{
					Value: "\"\"",
					Name:  name + "Empty",
				})
			} else {
				values = append(values, EnumValue{
					Value: "\"" + v + "\"",
					Name:  name + xstrings.ToCamelCase(v),
				})
			}
		case bool:
			strBool := strconv.FormatBool(v)
			values = append(values, EnumValue{
				Value: v,
				Name:  name + xstrings.ToCamelCase(strBool),
			})
		case float64:
			strFloat := strings.Replace(strconv.FormatFloat(v, 'g', -1, 64), "-", "negative", 1)
			values = append(values, EnumValue{
				Value: v,
				Name:  name + "Number" + xstrings.ToCamelCase(strFloat),
			})

		default:
			return fmt.Errorf("unhandled error type %T,%v", attrValue, attrValue)
		}
	}

	ConstOrVar := Const
	Type := "string"
	for _, attrValue := range attr.Values {
		if _, ok := attrValue.(bool); ok {
			ConstOrVar = Var
			Type = "interface{}"
			break
		}
		if _, ok := attrValue.(float64); ok {
			ConstOrVar = Var
			Type = "interface{}"
			break
		}
	}

	enum := Enum{
		Name:        name,
		Description: attr.Description,
		Values:      values,
		ConstOrVar:  ConstOrVar,
		Type:        Type,
	}

	r.Enums = append(r.Enums, enum)
	return nil
}

func (r *TraceFile) parseFlaglist(name string, attr *Attribute) error {

	flags := make([]FlaglistValue, 0, len(attr.Flags))
	for _, attrValue := range attr.Flags {
		if attrValue == "" {
			flags = append(flags, FlaglistValue{
				Value: "\"\"",
				Name:  name + "Empty",
			})
		} else {
			flags = append(flags, FlaglistValue{
				Value: "\"" + attrValue + "\"",
				Name:  name + xstrings.ToCamelCase(attrValue),
			})
		}
	}

	extra := make([]FlaglistValue, 0, len(attr.Extras))
	for _, attrValue := range attr.Extras {
		switch v := attrValue.(type) {
		case string:
			if v == "" {
				extra = append(extra, FlaglistValue{
					Value: "\"\"",
					Name:  name + "Empty",
				})
			} else {
				extra = append(extra, FlaglistValue{
					Value: "\"" + v + "\"",
					Name:  name + xstrings.ToCamelCase(v),
				})
			}
		case bool:
			strBool := strconv.FormatBool(v)
			extra = append(extra, FlaglistValue{
				Value: v,
				Name:  name + xstrings.ToCamelCase(strBool),
			})
		case float64:
			strFloat := strings.Replace(strconv.FormatFloat(v, 'g', -1, 64), "-", "negative", 1)
			extra = append(extra, FlaglistValue{
				Value: v,
				Name:  name + "Number" + xstrings.ToCamelCase(strFloat),
			})

		default:
			return fmt.Errorf("unhandled error type %T,%v", attrValue, attrValue)
		}
	}

	ConstOrVar := Const
	Type := "string"
	for _, attrValue := range attr.Values {
		if _, ok := attrValue.(bool); ok {
			ConstOrVar = Var
			Type = "interface{}"
			break
		}
		if _, ok := attrValue.(float64); ok {
			ConstOrVar = Var
			Type = "interface{}"
			break
		}
	}

	flaglist := FlagList{
		Name:        name,
		Description: attr.Description,
		ConstOrVar:  ConstOrVar,
		Type:        Type,
		Flags:       flags,
		Extra:       extra,
	}

	r.FlagLists = append(r.FlagLists, flaglist)
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
					Description: []string{"is the type of the plot"},
				},
			},
		},
		Objects:   []Struct{},
		Enums:     []Enum{},
		FlagLists: []FlagList{},
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
	for i := range traceFile.Enums {
		err := r.tmpl.ExecuteTemplate(w, "enum.tmpl", traceFile.Enums[i])
		if err != nil {
			return err
		}
	}
	for i := range traceFile.FlagLists {
		err := r.tmpl.ExecuteTemplate(w, "flaglist.tmpl", traceFile.FlagLists[i])
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
	Trace     Struct
	Objects   []Struct
	Enums     []Enum
	FlagLists []FlagList
}

type ConstOrVar string

const (
	Const ConstOrVar = "const"
	Var   ConstOrVar = "var"
)

type EnumValue struct {
	Name  string
	Value interface{}
}
type Enum struct {
	Name        string
	Description string
	Type        string
	ConstOrVar  ConstOrVar
	Values      []EnumValue
}

type FlaglistValue struct {
	Name  string
	Value interface{}
}

type FlagList struct {
	Name        string
	Description string
	Type        string
	ConstOrVar  ConstOrVar
	Flags       []FlaglistValue
	Extra       []FlaglistValue
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
	Description []string
}

var ValTypeMap = map[ValType]string{
	ValTypeInfoArray: "interface{}",
	ValTypeSubplotID: "String",
	ValTypeDataArray: "interface{}",
	ValTypeAny:       "interface{}",
	ValTypeAngle:     "float64",
	ValTypeColor:     "String",
	ValTypeNumber:    "float64",
	ValTypeString:    "String",
	ValTypeInteger:   "int64",
	ValTypeBoolean:   "Bool",
}
