package enums

import (
	"encoding/json"
	"text/template"
)

// CSharpType represents any C# scalar type, but also some reference types.
type CSharpType uint8

// C# scalar types.
const (
	CSharpTypeUnknown CSharpType = iota
	CSharpTypeBool
	CSharpTypeByte
	CSharpTypeFloat
	CSharpTypeDouble
	CSharpTypeDecimal
	CSharpTypeInt
	CSharpTypeLong

	CSharpTypeDateTime
	CSharpTypeDateTimeOffset
	CSharpTypeTimeSpan

	CSharpTypeString

	CSharpTypeByteArray
	CSharpTypeGuid
)

// CSharpTypeFromName returns the CSharpType for the given name.
func CSharpTypeFromName(name string) CSharpType {
	switch name {
	case "bool":
		return CSharpTypeBool
	case "byte":
		return CSharpTypeByte
	case "float":
		return CSharpTypeFloat
	case "double":
		return CSharpTypeDouble
	case "decimal":
		return CSharpTypeDecimal
	case "int":
		return CSharpTypeInt
	case "long":
		return CSharpTypeLong
	case "DateTime":
		return CSharpTypeDateTime
	case "DateTimeOffset":
		return CSharpTypeDateTimeOffset
	case "TimeSpan":
		return CSharpTypeTimeSpan
	case "string":
		return CSharpTypeString
	case "byte_array":
		return CSharpTypeByteArray
	case "Guid":
		return CSharpTypeGuid
	default:
		return CSharpTypeUnknown
	}
}

// Name returns the name of the CSharpType.
func (t CSharpType) Name() string {
	switch t {
	case CSharpTypeBool:
		return "bool"
	case CSharpTypeByte:
		return "byte"
	case CSharpTypeFloat:
		return "float"
	case CSharpTypeDouble:
		return "double"
	case CSharpTypeDecimal:
		return "decimal"
	case CSharpTypeInt:
		return "int"
	case CSharpTypeLong:
		return "long"
	case CSharpTypeDateTime:
		return "DateTime"
	case CSharpTypeDateTimeOffset:
		return "DateTimeOffset"
	case CSharpTypeTimeSpan:
		return "TimeSpan"
	case CSharpTypeString:
		return "string"
	case CSharpTypeByteArray:
		return "byte[]"
	case CSharpTypeGuid:
		return "Guid"
	default:
		return "Unknown"
	}
}

// CSharpTypeFuncMap is a map of CSharpType to a function.
// Used in templates to generate code.
var CSharpTypeFuncMap = template.FuncMap{
	"bool": func() CSharpType {
		return CSharpTypeBool
	},
	"byte": func() CSharpType {
		return CSharpTypeByte
	},
	"float": func() CSharpType {
		return CSharpTypeFloat
	},
	"double": func() CSharpType {
		return CSharpTypeDouble
	},
	"decimal": func() CSharpType {
		return CSharpTypeDecimal
	},
	"int": func() CSharpType {
		return CSharpTypeInt
	},
	"long": func() CSharpType {
		return CSharpTypeLong
	},
	"DateTime": func() CSharpType {
		return CSharpTypeDateTime
	},
	"DateTimeOffset": func() CSharpType {
		return CSharpTypeDateTimeOffset
	},
	"TimeSpan": func() CSharpType {
		return CSharpTypeTimeSpan
	},
	"string": func() CSharpType {
		return CSharpTypeString
	},
	"byte_array": func() CSharpType {
		return CSharpTypeByteArray
	},
	"Guid": func() CSharpType {
		return CSharpTypeGuid
	},
}

// MarshalYAML marshals the CSharpType to YAML.
func (t CSharpType) MarshalYAML() (interface{}, error) {
	return t.Name(), nil
}

// UnmarshalYAML unmarshals the CSharpType from YAML.
func (t *CSharpType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	if err := unmarshal(&name); err != nil {
		return err
	}
	*t = CSharpTypeFromName(name)
	return nil
}

// MarshalJSON marshals the CSharpType to JSON.
func (t CSharpType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name())
}

// UnmarshalJSON unmarshals the CSharpType from JSON.
func (t *CSharpType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = CSharpTypeFromName(name)
	return nil
}
