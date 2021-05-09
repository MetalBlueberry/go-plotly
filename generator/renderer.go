package generator

import (
	"embed"
	"fmt"
	"io"
	"path"
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

// parseAttributes returns a []StructField containing all the fields for the attributes parsed.
// Nested structures found such as enums, flags u other objects are stored in the TypeFile caller.
func (r *TypeFile) parseAttributes(namePrefix string, typePrefix string, attr map[string]*Attribute) ([]StructField, error) {
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
			name := namePrefix + xstrings.ToCamelCase(attr.Name)
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
			name := namePrefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseFlaglist(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse flaglist %s, %w", name, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        typePrefix + xstrings.ToCamelCase(attr.Name),
				Description: []string{string(attr.ValType) + " " + attr.Description},
			})

		case attr.ValType == ValTypeEnum:
			typeName := typePrefix + xstrings.ToCamelCase(attr.Name)
			valueName := namePrefix + xstrings.ToCamelCase(attr.Name)
			err := r.parseEnum(typeName, valueName, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse enum %s, %w", typeName, err)
			}
			fields = append(fields, StructField{
				Name:        xstrings.ToCamelCase(attr.Name),
				JSONName:    attr.Name,
				Type:        typeName,
				Description: []string{string(attr.ValType) + " " + attr.Description},
			})

		case attr.ValType == ValTypeColorscale:
			fields = append(fields, StructField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     "ColorScale",
				Description: []string{
					string(attr.ValType),
					attr.Description,
					fmt.Sprintf("Default: %v", attr.Dflt),
				},
			})

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

func (r *TypeFile) parseObject(name string, attr *Attribute) error {
	objStruct := Struct{
		Name:        name,
		Description: attr.Description,
		Fields:      []StructField{},
	}

	fields, err := r.parseAttributes(objStruct.Name, objStruct.Name, attr.Attributes)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	objStruct.Fields = fields

	r.Objects = append(r.Objects, objStruct)
	return nil
}

func (r *TypeFile) parseEnum(typeName string, valuePrefix string, attr *Attribute) error {

	values := make([]EnumValue, 0, len(attr.Values))
	for _, attrValue := range attr.Values {
		switch v := attrValue.(type) {
		case string:
			if v == "" {
				values = append(values, EnumValue{
					Value: "\"\"",
					Name:  valuePrefix + "Empty",
				})
			} else {
				values = append(values, EnumValue{
					Value: "\"" + cleanValue(v) + "\"",
					Name:  valuePrefix + xstrings.ToCamelCase(v),
				})
			}
		case bool:
			strBool := strconv.FormatBool(v)
			values = append(values, EnumValue{
				Value: v,
				Name:  valuePrefix + xstrings.ToCamelCase(strBool),
			})
		case float64:
			strFloat := strings.Replace(strconv.FormatFloat(v, 'g', -1, 64), "-", "negative", 1)
			values = append(values, EnumValue{
				Value: v,
				Name:  valuePrefix + "Number" + xstrings.ToCamelCase(strFloat),
			})

		default:
			return fmt.Errorf("unhandled error type %T,%v", attrValue, attrValue)
		}
	}

	for i := range values {
		values[i].Name = cleanName(values[i].Name)
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

	duplicated := map[string]int{}
	for i := range values {
		_, ok := duplicated[values[i].Name]
		if !ok {
			duplicated[values[i].Name] = i
			continue
		}
		values[duplicated[values[i].Name]].Name += "1"
		index := 2
		for {
			values[i].Name = values[i].Name + strconv.Itoa(index)
			_, ok := duplicated[values[i].Name]
			if !ok {
				duplicated[values[i].Name] = i
				break
			}
			index++
		}
	}

	enum := Enum{
		Name:        typeName,
		Description: attr.Description,
		Values:      values,
		ConstOrVar:  ConstOrVar,
		Type:        Type,
	}

	r.Enums = append(r.Enums, enum)
	return nil
}

func (r *TypeFile) parseFlaglist(name string, attr *Attribute) error {

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

	traceFile := TypeFile{
		MainType: Struct{
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

	fields, err := traceFile.parseAttributes(traceFile.MainType.Name, traceFile.MainType.Name, trace.Attributes.Names)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	traceFile.MainType.Fields = append(traceFile.MainType.Fields, fields...)

	fmt.Fprintf(w, `package grob

var TraceType%s TraceType = "%s"

func (trace *%s) GetType() TraceType {
	return TraceType%s
}
`,
		traceFile.MainType.Name,
		traceName,
		traceFile.MainType.Name,
		traceFile.MainType.Name,
	)

	err = r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceFile.MainType)
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

func (r *Renderer) WriteTraces(dir string) error {
	traceNames := make([]string, 0, len(r.root.Schema.Traces))
	for n := range r.root.Schema.Traces {
		traceNames = append(traceNames, n)
	}
	sort.Strings(traceNames)
	for _, name := range traceNames {
		f, err := r.fs.Create(path.Join(dir, name+"_gen.go"))
		if err != nil {
			return fmt.Errorf("cannot create file, %w", err)
		}
		defer f.Close()
		err = r.WriteTrace(name, f)
		if err != nil {
			return fmt.Errorf("cannot write trace, %w", err)
		}
	}
	return nil
}

func (r *Renderer) WriteLayout(dir string) error {
	w, err := r.fs.Create(path.Join(dir, "layout_gen.go"))
	if err != nil {
		return err
	}
	defer w.Close()

	traceFile := TypeFile{
		MainType: Struct{
			Name:        "Layout",
			Description: "Plot layout options",
			Fields:      []StructField{},
		},
		Objects:   []Struct{},
		Enums:     []Enum{},
		FlagLists: []FlagList{},
	}

	fields, err := traceFile.parseAttributes(traceFile.MainType.Name, traceFile.MainType.Name, r.root.Schema.Layout.LayoutAttributes.Names)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	traceFile.MainType.Fields = append(traceFile.MainType.Fields, fields...)

	for name, trace := range r.root.Schema.Traces {
		fields, err := traceFile.parseAttributes(xstrings.ToCamelCase(name), "Layout", trace.LayoutAttributes.Names)
		if err != nil {
			return fmt.Errorf("cannot parse attributes, %w", err)
		}
		traceFile.MainType.Fields = append(traceFile.MainType.Fields, fields...)
	}

	// remove duplicate fields
	uniqueFields := make([]StructField, 0, len(traceFile.MainType.Fields))
	fieldMap := map[string]int{}
	for i, field := range traceFile.MainType.Fields {
		_, ok := fieldMap[field.Name]
		if !ok {
			fieldMap[field.Name] = i
			uniqueFields = append(uniqueFields, field)
			continue
		}
	}
	traceFile.MainType.Fields = uniqueFields

	// merge duplicate enums
	uniqueEnums := make([]Enum, 0, len(traceFile.Enums))
	enumMap := map[string]int{}
	for _, enum := range traceFile.Enums {
		previous, ok := enumMap[enum.Name]
		if !ok {
			uniqueEnums = append(uniqueEnums, enum)
			enumMap[enum.Name] = len(uniqueEnums) - 1
			continue
		}
		uniqueEnums[previous].Values = append(uniqueEnums[previous].Values, enum.Values...)
	}
	traceFile.Enums = uniqueEnums

	fmt.Fprint(w, `package grob

`)

	err = r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceFile.MainType)
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

func (r *Renderer) WriteConfig(dir string) error {
	w, err := r.fs.Create(path.Join(dir, "config_gen.go"))
	if err != nil {
		return err
	}
	defer w.Close()

	traceFile := TypeFile{
		MainType: Struct{
			Name:        "Config",
			Description: "Plot config options",
			Fields:      []StructField{},
		},
		Objects:   []Struct{},
		Enums:     []Enum{},
		FlagLists: []FlagList{},
	}
	fields, err := traceFile.parseAttributes(traceFile.MainType.Name, traceFile.MainType.Name, r.root.Schema.Layout.LayoutAttributes.Names)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	traceFile.MainType.Fields = append(traceFile.MainType.Fields, fields...)

	fmt.Fprint(w, `package grob

`)

	err = r.tmpl.ExecuteTemplate(w, "trace.tmpl", traceFile.MainType)
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

type TypeFile struct {
	MainType  Struct
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
	ValTypeDataArray:  "interface{}",
	ValTypeEnum:       "NO-TYPE",
	ValTypeBoolean:    "Bool",
	ValTypeNumber:     "float64",
	ValTypeInteger:    "int64",
	ValTypeString:     "String",
	ValTypeColor:      "Color",
	ValTypeColorlist:  "ColorList",
	ValTypeColorscale: "ColorScale",
	ValTypeAngle:      "float64",
	ValTypeSubplotID:  "String",
	ValTypeFlagList:   "NO-TYPE",
	ValTypeAny:        "interface{}",
	ValTypeInfoArray:  "interface{}",
}

var SymbolMap = []string{
	"=", "Eq",
	">", "Gt",
	"-", "Hyphen",
	"<", "Lt",
	"|", "Or",
	"/", "Slash",
	"\\", "Doublebackslash",
	"^", "Cape",
	"(", "Lpar",
	")", "Rpar",
	"[", "Lbracket",
	"]", "Rbracket",
	"+", "Plus",
	"?", "Question",
	"$", "Dollar",
}

func cleanName(name string) string {
	replacer := strings.NewReplacer(SymbolMap...)
	return replacer.Replace(name)
}
func cleanValue(value string) string {
	return strings.ReplaceAll(value, "\\", "\\\\")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
