package generator

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/huandu/xstrings"
)

type typeFile struct {
	MainType  sstruct
	Objects   []sstruct
	Enums     enumFields
	FlagLists []flagList
}

type sstruct struct {
	Name        string
	Description string
	Fields      structFields
}

type structFields []structField

func (a structFields) Len() int           { return len(a) }
func (a structFields) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a structFields) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type structField struct {
	Name        string
	JSONName    string
	Type        string
	Description []string
}

type enumFields []enumFile

func (a enumFields) Len() int           { return len(a) }
func (a enumFields) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a enumFields) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type enumFile struct {
	Name        string
	Description string
	Type        string
	ConstOrVar  costOrVar
	Values      []enumValue
}

type enumValue struct {
	Name  string
	Value interface{}
}

type flagList struct {
	Name        string
	Description string
	Type        string
	ConstOrVar  costOrVar
	Flags       []flagListValue
	Extra       []flagListValue
}

type flagListValue struct {
	Name  string
	Value interface{}
}

type costOrVar string

const (
	constant costOrVar = "const"
	variable costOrVar = "var"
)

// parseAttributes returns a []StructField containing all the fields for the attributes parsed.
// Nested structures found such as enums, flags u other objects are stored in the TypeFile caller.
func (file *typeFile) parseAttributes(namePrefix string, typePrefix string, attr map[string]*Attribute) ([]structField, error) {
	fields := make([]structField, 0, len(attr))

	for _, name := range sortKeys(attr) {
		if name == "_deprecated" {
			continue
		}

		attr := attr[name]
		// if arrays are allowed, example marker sizes, then javascript allows multiple different inputs hard to map in golang:
		// [100,50,25,25,50,100] -> use different sizes for the markers
		// [100] -> only the first marker will have size 100, all others have 0 size
		// 100 -> all markers have size 100
		// this can only be captures by using interface and allowing the user to define the used type themselves.
		// this sadly means that there is no more type checking on those entries
		var adjust = func(typename string) string { return typename }
		if attr.ArrayOK {
			adjust = func(typename string) string { return "interface{}" }
		}

		switch {
		case attr.Role == RoleObject && len(attr.Items) > 0:
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust("interface{}"),
				Description: []string{
					"It's an items array and what goes inside it's... messy... check the docs",
					"I will be happy if you want to contribute by implementing this",
					"just raise an issue before you start so we do not overlap",
				},
			})

		case attr.Role == RoleObject:
			name := namePrefix + xstrings.ToCamelCase(attr.Name)
			err := file.parseObject(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse object %s, %w", name, err)
			}
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust("*" + name),
				Description: []string{
					"role: Object",
				},
			})

		case attr.ValType == ValTypeFlagList:
			name := namePrefix + xstrings.ToCamelCase(attr.Name)
			err := file.parseFlaglist(name, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse flaglist %s, %w", name, err)
			}
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust(typePrefix + xstrings.ToCamelCase(attr.Name)),
				Description: []string{
					fmt.Sprintf("default: %s", attr.Dflt),
					fmt.Sprintf("type: %s", attr.ValType),
					attr.Description,
				},
			})

		case attr.ValType == ValTypeEnum:
			typeName := typePrefix + xstrings.ToCamelCase(attr.Name)
			valueName := namePrefix + xstrings.ToCamelCase(attr.Name)
			err := file.parseEnum(typeName, valueName, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot parse enum %s, %w", typeName, err)
			}
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust(typeName),
				Description: []string{
					fmt.Sprintf("default: %s", attr.Dflt),
					fmt.Sprintf("type: %s", attr.ValType),
					attr.Description,
				},
			})

		case attr.ValType == ValTypeColorscale:
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust("ColorScale"),
				Description: []string{
					fmt.Sprintf("default: %s", attr.Dflt),
					fmt.Sprintf("type: %s", attr.ValType),
					attr.Description,
				},
			})

		default:
			ty := valTypeMap[attr.ValType]
			fields = append(fields, structField{
				Name:     xstrings.ToCamelCase(attr.Name),
				JSONName: attr.Name,
				Type:     adjust(ty),
				Description: []string{
					fmt.Sprintf("arrayOK: %t", attr.ArrayOK),
					fmt.Sprintf("type: %s", attr.ValType),
					attr.Description,
				},
			})
		}
	}
	return fields, nil
}

func (file *typeFile) parseObject(name string, attr *Attribute) error {
	objStruct := sstruct{
		Name:        name,
		Description: attr.Description,
		Fields:      []structField{},
	}

	fields, err := file.parseAttributes(objStruct.Name, objStruct.Name, attr.Attributes)
	if err != nil {
		return fmt.Errorf("cannot parse attributes, %w", err)
	}
	objStruct.Fields = fields

	file.Objects = append(file.Objects, objStruct)
	return nil
}

func (file *typeFile) parseEnum(typeName string, valuePrefix string, attr *Attribute) error {

	values := make([]enumValue, 0, len(attr.Values))
	for _, attrValue := range attr.Values {
		switch v := attrValue.(type) {
		case string:
			if v == "" {
				values = append(values, enumValue{
					Value: "\"\"",
					Name:  valuePrefix + "Empty",
				})
			} else {
				values = append(values, enumValue{
					Value: "\"" + cleanValue(v) + "\"",
					Name:  valuePrefix + xstrings.ToCamelCase(v),
				})
			}
		case bool:
			strBool := strconv.FormatBool(v)
			values = append(values, enumValue{
				Value: v,
				Name:  valuePrefix + xstrings.ToCamelCase(strBool),
			})
		case float64:
			strFloat := strings.Replace(strconv.FormatFloat(v, 'g', -1, 64), "-", "negative", 1)
			values = append(values, enumValue{
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

	ConstOrVar := constant
	Type := "string"
	for _, attrValue := range attr.Values {
		if _, ok := attrValue.(bool); ok {
			ConstOrVar = variable
			Type = "interface{}"
			break
		}
		if _, ok := attrValue.(float64); ok {
			ConstOrVar = variable
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

	enum := enumFile{
		Name:        typeName,
		Description: attr.Description,
		Values:      values,
		ConstOrVar:  ConstOrVar,
		Type:        Type,
	}

	file.Enums = append(file.Enums, enum)
	return nil
}

func (file *typeFile) parseFlaglist(name string, attr *Attribute) error {

	flags := make([]flagListValue, 0, len(attr.Flags))
	for _, attrValue := range attr.Flags {
		if attrValue == "" {
			flags = append(flags, flagListValue{
				Value: "\"\"",
				Name:  name + "Empty",
			})
		} else {
			flags = append(flags, flagListValue{
				Value: "\"" + attrValue + "\"",
				Name:  name + xstrings.ToCamelCase(attrValue),
			})
		}
	}

	extra := make([]flagListValue, 0, len(attr.Extras))
	for _, attrValue := range attr.Extras {
		switch v := attrValue.(type) {
		case string:
			if v == "" {
				extra = append(extra, flagListValue{
					Value: "\"\"",
					Name:  name + "Empty",
				})
			} else {
				extra = append(extra, flagListValue{
					Value: "\"" + v + "\"",
					Name:  name + xstrings.ToCamelCase(v),
				})
			}
		case bool:
			strBool := strconv.FormatBool(v)
			extra = append(extra, flagListValue{
				Value: v,
				Name:  name + xstrings.ToCamelCase(strBool),
			})
		case float64:
			strFloat := strings.Replace(strconv.FormatFloat(v, 'g', -1, 64), "-", "negative", 1)
			extra = append(extra, flagListValue{
				Value: v,
				Name:  name + "Number" + xstrings.ToCamelCase(strFloat),
			})

		default:
			return fmt.Errorf("unhandled error type %T,%v", attrValue, attrValue)
		}
	}

	ConstOrVar := constant
	Type := "string"
	for _, attrValue := range attr.Extras {
		if _, ok := attrValue.(bool); ok {
			ConstOrVar = variable
			Type = "interface{}"
			break
		}
		if _, ok := attrValue.(float64); ok {
			ConstOrVar = variable
			Type = "interface{}"
			break
		}
	}

	flaglist := flagList{
		Name:        name,
		Description: attr.Description,
		ConstOrVar:  ConstOrVar,
		Type:        Type,
		Flags:       flags,
		Extra:       extra,
	}

	file.FlagLists = append(file.FlagLists, flaglist)
	return nil
}

func sortKeys[T any](attr map[string]*T) []string {
	keys := make([]string, 0, len(attr))
	for k := range attr {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
