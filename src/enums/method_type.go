package enums

import (
	"encoding/json"
	"text/template"
)

// MethodType represents the type of a method. i.e. Delete, Get, Insert, Update etc.
type MethodType uint8

// Method types.
const (
	MethodTypeUnknown MethodType = iota
	MethodTypeMultiGet
	MethodTypeGetOrCreate
	MethodTypeGetCollection
	MethodTypeGetCollectionPaged
	MethodTypeGetCollectionExclusive
	MethodTypeGetCount
	MethodTypeLookup
	MethodTypeMustGet
)

// MethodTypeFromName returns the MethodType for the given name.
func MethodTypeFromName(name string) MethodType {
	switch name {
	case "MultiGet":
		return MethodTypeMultiGet
	case "GetOrCreate":
		return MethodTypeGetOrCreate
	case "GetCollection":
		return MethodTypeGetCollection
	case "GetCollectionPaged":
		return MethodTypeGetCollectionPaged
	case "GetCollectionExclusive":
		return MethodTypeGetCollectionExclusive
	case "GetCount":
		return MethodTypeGetCount
	case "Lookup":
		return MethodTypeLookup
	case "MustGet":
		return MethodTypeMustGet
	default:
		return MethodTypeUnknown
	}
}

// Name returns the name of the MethodType.
func (t MethodType) Name() string {
	switch t {
	case MethodTypeMultiGet:
		return "MultiGet"
	case MethodTypeGetOrCreate:
		return "GetOrCreate"
	case MethodTypeGetCollection:
		return "GetCollection"
	case MethodTypeGetCollectionPaged:
		return "GetCollectionPaged"
	case MethodTypeGetCollectionExclusive:
		return "GetCollectionExclusive"
	case MethodTypeGetCount:
		return "GetCount"
	case MethodTypeLookup:
		return "Lookup"
	case MethodTypeMustGet:
		return "MustGet"
	default:
		return "Unknown"
	}
}

// MethodTypeFuncMap is a map of MethodType to a function.
// Used in templates to generate code.
var MethodTypeFuncMap = template.FuncMap{
	"MultiGet": func() MethodType {
		return MethodTypeMultiGet
	},
	"GetOrCreate": func() MethodType {
		return MethodTypeGetOrCreate
	},
	"GetCollection": func() MethodType {
		return MethodTypeGetCollection
	},
	"GetCollectionPaged": func() MethodType {
		return MethodTypeGetCollectionPaged
	},
	"GetCollectionExclusive": func() MethodType {
		return MethodTypeGetCollectionExclusive
	},
	"GetCount": func() MethodType {
		return MethodTypeGetCount
	},
	"Lookup": func() MethodType {
		return MethodTypeLookup
	},
	"MustGet": func() MethodType {
		return MethodTypeMustGet
	},
}

// MarshalYAML marshals the MethodType to YAML.
func (t MethodType) MarshalYAML() (interface{}, error) {
	return t.Name(), nil
}

// UnmarshalYAML unmarshals the MethodType from YAML.
func (t *MethodType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	if err := unmarshal(&name); err != nil {
		return err
	}

	*t = MethodTypeFromName(name)
	return nil
}

// MarshalJSON marshals the MethodType to JSON.
func (t MethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name())
}

// UnmarshalJSON unmarshals the MethodType from JSON.
func (t *MethodType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	*t = MethodTypeFromName(name)
	return nil
}
