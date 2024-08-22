package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

// LoadSchema loads the plotly schema file
func LoadSchema(r io.Reader) (*Root, error) {
	decoder := json.NewDecoder(r)
	root := &Root{}
	err := decoder.Decode(root)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Root represents the root of the plotly schema
// It is basically the struct representation of the json schema content
type Root struct {
	SHA1     string  `json:"sha1,omitempty"`
	Modified bool    `json:"modified,omitempty"`
	Schema   *Schema `json:"schema,omitempty"`
}

type Schema struct {
	Defs   *Defs  `json:"defs,omitempty"`
	Traces Traces `json:"traces,omitempty"`
	Layout Layout `json:"layout,omitempty"`
	// Transforms *Transforms `json:"transforms,omitempty"`
	Frames    *Frames              `json:"frames,omitempty"`
	Animation *AnimationAttributes `json:"animation,omitempty"`
	Config    *ConfigAttributes    `json:"config,omitempty"`
}

type Frames struct {
	Attribute *Attribute `json:"-"`
}

type AnimationAttributes struct {
	Names map[string]*Attribute `json:"-"`
}

type ConfigAttributes struct {
	Names map[string]*Attribute `json:"-"`
}

type Layout struct {
	LayoutAttributes LayoutAttributes `json:"layoutAttributes,omitempty"`
}

type Defs struct {
	ValObjects map[string]*ValObject `json:"valObjects,omitempty"`
	MetaKeys   []string              `json:"meta_keys,omitempty"`
	// EditType     *EditType
	// ImpliedEdits *ImpliedEdits
}

type ValObject struct {
	Description  string   `json:"description,omitempty"`
	RequiredOpts []string `json:"requiredOpts,omitempty"`
	OtherOpts    []string `json:"otherOpts,omitempty"`
}

type Traces map[string]*Trace

type Trace struct {
	Meta *Meta `json:"meta,omitempty"`
	// Categories []string   `json:"categories,omitempty"`
	Animatable       bool             `json:"animatable,omitempty"`
	Type             string           `json:"type,omitempty"`
	Attributes       TraceAttributes  `json:"attributes,omitempty"`
	LayoutAttributes LayoutAttributes `json:"layoutAttributes,omitempty"`
}

type Meta struct {
	Description string `json:"description,omitempty"`
}

type TraceAttributes struct {
	Type  string                `json:"type"`
	Names map[string]*Attribute `json:"-"`
}

func (attr *TraceAttributes) Sorted() []string {
	keys := make([]string, 0, len(attr.Names))
	for k := range attr.Names {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (attr *TraceAttributes) UnmarshalJSON(b []byte) error {
	var err error

	fields := map[string]json.RawMessage{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fields["type"], &attr.Type)
	if err != nil {
		return err
	}
	delete(fields, "type")

	names, err := parseFields(fields, nil)
	if err != nil {
		return err
	}
	attr.Names = names

	return nil
}

type LayoutAttributes struct {
	Names map[string]*Attribute `json:"-"`
}

func (attr *LayoutAttributes) UnmarshalJSON(b []byte) error {
	var err error

	fields := map[string]json.RawMessage{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	names, err := parseFields(fields, nil)
	if err != nil {
		return err
	}
	attr.Names = names
	return nil
}

func (attr *Frames) UnmarshalJSON(b []byte) error {
	var err error

	field := json.RawMessage{}
	err = json.Unmarshal(b, &field)
	if err != nil {
		return err
	}

	items, err := parseField("items", field, nil)
	if err != nil {
		return err
	}
	attr.Attribute = items
	return nil
}

func (attr *AnimationAttributes) UnmarshalJSON(b []byte) error {
	var err error

	fields := map[string]json.RawMessage{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	names, err := parseFields(fields, nil)
	if err != nil {
		return err
	}
	attr.Names = names
	return nil
}

func (attr *ConfigAttributes) UnmarshalJSON(b []byte) error {
	var err error

	fields := map[string]json.RawMessage{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	names, err := parseFields(fields, nil)
	if err != nil {
		return err
	}
	attr.Names = names
	return nil
}

func parseField(name string, value json.RawMessage, parent *Attribute) (_ *Attribute, err error) {
	role := &struct {
		Role  Role            `json:"role,omitempty"`
		Items json.RawMessage `json:"items,omitempty"`
	}{}
	err = json.Unmarshal(value, role)
	if err != nil {
		log.Printf("cannot detect role for attribute %s on %v, %s", name, parent, err)
		return
	}
	switch {
	case role.Role == RoleObject && len(role.Items) > 0:
		subFields := map[string]json.RawMessage{}

		err = json.Unmarshal(role.Items, &subFields)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal attribute subfields, %s, %w", name, err)
		}

		attr := &Attribute{
			Role:   role.Role,
			Name:   name,
			Parent: parent,
		}
		subAttr, err := parseFields(subFields, attr)
		if err != nil {
			return nil, fmt.Errorf("on %s, %w", name, err)
		}
		attr.Items = subAttr
		return attr, nil

	case role.Role == RoleObject:
		subFields := map[string]json.RawMessage{}

		err = json.Unmarshal(value, &subFields)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal attribute subfields, %s, %w", name, err)
		}

		attr := &Attribute{
			Name:   name,
			Parent: parent,
		}

		err := UnmarshalRole(subFields["role"], &attr.Role)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal role, %w", err)
		}

		if editType, ok := subFields["editType"]; ok {
			err = json.Unmarshal(editType, &attr.EditType)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal editType, %w", err)
			}
		}

		if description, ok := subFields["description"]; ok {
			err = json.Unmarshal(description, &attr.Description)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal description, %w", err)
			}
		}

		delete(subFields, "role")
		delete(subFields, "editType")
		delete(subFields, "description")

		subAttr, err := parseFields(subFields, attr)
		if err != nil {
			return nil, fmt.Errorf("on %s, %w", name, err)
		}

		attr.Attributes = subAttr
		return attr, nil
	default:
		attr := &Attribute{
			Name:   name,
			Parent: parent,
		}
		err = json.Unmarshal(value, attr)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal attribute, %s, %w", name, err)
		}
		return attr, nil
	}
}

func parseFields(fields map[string]json.RawMessage, parent *Attribute) (_ map[string]*Attribute, err error) {
	attributes := make(map[string]*Attribute)
	for name, value := range fields {
		if isMetaKey(name) {
			// is a metakey and it is not required to
			// generate the schema
			continue
		}

		attr, err := parseField(name, value, parent)
		if err != nil {
			return nil, fmt.Errorf("cannot parse field %s, %w", name, err)
		}
		attributes[name] = attr
	}
	return attributes, nil
}

type Role string

const (
	RoleObject Role = "object"
	RoleInfo   Role = "info"
	RoleStyle  Role = "style"
	RoleData   Role = "data"
)

type ValType string

const (
	ValTypeDataArray  ValType = "data_array"
	ValTypeEnum       ValType = "enumerated"
	ValTypeBoolean    ValType = "boolean"
	ValTypeNumber     ValType = "number"
	ValTypeInteger    ValType = "integer"
	ValTypeString     ValType = "string"
	ValTypeColor      ValType = "color"
	ValTypeColorlist  ValType = "colorlist"
	ValTypeColorscale ValType = "colorscale"
	ValTypeAngle      ValType = "angle"
	ValTypeSubplotID  ValType = "subplotid"
	ValTypeFlagList   ValType = "flaglist"
	ValTypeAny        ValType = "any"
	ValTypeInfoArray  ValType = "info_array"
)

func UnmarshalRole(obj json.RawMessage, role *Role) error {
	err := json.Unmarshal(obj, role)
	if err != nil {
		return err
	}

	switch *role {
	case RoleObject, RoleInfo, RoleStyle, RoleData:
		return nil
	default:
		return fmt.Errorf("Invalid role %s", *role)
	}
}

type Attribute struct {
	Role        Role   `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
	EditType    string `json:"editType,omitempty"`

	ValType ValType       `json:"valType,omitempty"`
	Values  []interface{} `json:"values,omitempty"`
	Flags   []string      `json:"flags,omitempty"`
	Extras  []interface{} `json:"extras,omitempty"`
	Dflt    interface{}   `json:"dflt,omitempty"`
	Min     json.Number   `json:"min,omitempty"`
	Max     json.Number   `json:"max,omitempty"`
	ArrayOK bool          `json:"arrayOk,omitempty"`
	Anim    bool          `json:"anim,omitempty"`

	Name       string                `json:"-"`
	Attributes map[string]*Attribute `json:"-"`
	Items      map[string]*Attribute `json:"-"`
	Parent     *Attribute            `json:"-"`
}

func (attr *Attribute) String() string {
	path := []string{}
	i := attr
	for i != nil {
		path = append(path, i.Name)
		i = i.Parent
	}
	return strings.Join(path, "->")
}

type MetaKey string

var MetaKeys = []MetaKey{
	"_isSubplotObj",
	"_isLinkedToArray",
	"_arrayAttrRegexps",
	"_deprecated",
	"description",
	"role",
	"editType",
	"impliedEdits",
}

func isMetaKey(v string) bool {
	for _, k := range MetaKeys {
		if string(k) == v {
			return true
		}
	}
	return false
}
