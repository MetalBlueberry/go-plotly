package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

func LoadSchema(r io.Reader) (*Root, error) {
	decoder := json.NewDecoder(r)
	root := &Root{}
	err := decoder.Decode(root)
	if err != nil {
		return nil, err
	}
	return root, nil
}

type Root struct {
	SHA1     string  `json:"sha1,omitempty"`
	Modified bool    `json:"modified,omitempty"`
	Schema   *Schema `json:"schema,omitempty"`
}

type Schema struct {
	Defs   *Defs  `json:"defs,omitempty"`
	Traces Traces `json:"traces,omitempty"`
	// Layout     *Layout     `json:"layout,omitempty"`
	// Transforms *Transforms `json:"transforms,omitempty"`
	// Frames     *Frames     `json:"frames,omitempty"`
	// Animation  *Animation  `json:"animation,omitempty"`
	// Config     *Config     `json:"config,omitempty"`
}

type Defs struct {
	ValObjects ValObjects `json:"valObjects,omitempty"`
	// MetaKeys   []string
	// EditType     *EditType
	// ImpliedEdits *ImpliedEdits
}

type ValObjects map[string]*ValObject

type ValObject struct {
	Description  string   `json:"description,omitempty"`
	RequiredOpts []string `json:"requiredOpts,omitempty"`
	OtherOpts    []string `json:"otherOpts,omitempty"`
}

type Traces map[string]*Trace

func (t Traces) Sorted() []string {
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type Trace struct {
	Meta *Meta `json:"meta,omitempty"`
	// Categories []string   `json:"categories,omitempty"`
	Animatable       bool                  `json:"animatable,omitempty"`
	Type             string                `json:"type,omitempty"`
	Attributes       Attributes            `json:"attributes,omitempty"`
	LayoutAttributes map[string]*Attribute `json:"layoutAttributes,omitempty"`
}

type Meta struct {
	Description string `json:"description,omitempty"`
}

type Attributes struct {
	Type  string                `json:"type"`
	Names map[string]*Attribute `json:"-"`
}

func (a *Attributes) Sorted() []string {
	keys := make([]string, 0, len(a.Names))
	for k := range a.Names {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (obj *Attributes) UnmarshalJSON(b []byte) error {
	var err error

	fields := map[string]json.RawMessage{}
	err = json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fields["type"], &obj.Type)
	if err != nil {
		return err
	}
	delete(fields, "type")

	names, err := parseFields(fields, nil)
	if err != nil {
		return err
	}
	obj.Names = names

	return nil
}

func parseFields(fields map[string]json.RawMessage, parent *Attribute) (_ map[string]*Attribute, err error) {
	attributes := make(map[string]*Attribute)
	for name, value := range fields {
		role := &struct {
			Role Role
		}{}
		err = json.Unmarshal(value, role)
		if err != nil {
			log.Printf("cannot detect role for attribute %s on %v, %s", name, parent, err)
			log.Printf("%s", value)
			continue
		}
		switch role.Role {
		case RoleObject:
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
			attributes[name] = attr
		default:
			attr := &Attribute{
				Name:   name,
				Parent: parent,
			}
			err = json.Unmarshal(value, attr)
			if err != nil {
				return nil, fmt.Errorf("cannot unmarshal attribute, %s, %w", name, err)
			}
			attributes[name] = attr
		}
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

	ValType string        `json:"valType,omitempty"`
	Values  []interface{} `json:"values,omitempty"`
	Flags   []string      `json:"flags,omitempty"`
	Extras  []string      `json:"extras,omitempty"`
	Dflt    interface{}   `json:"dflt,omitempty"`
	Min     json.Number   `json:"min,omitempty"`
	Max     json.Number   `json:"max,omitempty"`
	ArrayOK bool          `json:"arrayOk,omitempty"`
	Anim    bool          `json:"anim,omitempty"`

	Name       string                `json:"-"`
	Attributes map[string]*Attribute `json:"-"`
	Parent     *Attribute            `json:"-"`
}

func (attr *Attribute) SortedAttributes() []string {
	keys := make([]string, 0, len(attr.Attributes))
	for k := range attr.Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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
