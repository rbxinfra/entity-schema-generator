package enums

import (
	"encoding/json"
	"text/template"
)

// SqlDbType represents a SQL Server database type.
type SqlDbType uint8

// SQL Server database types.
const (
	SqlDbTypeUnknown SqlDbType = iota

	SqlDbTypeBit
	SqlDbTypeTinyInt
	SqlDbTypeFloat
	SqlDbTypeDouble
	SqlDbTypeDecimal
	SqlDbTypeInt
	SqlDbTypeBigInt

	SqlDbTypeDateTime
	SqlDbTypeDateTimeOffset
	SqlDbTypeTime

	SqlDbTypeVarChar
	SqlDbTypeNVarChar

	SqlDbTypeBinary
	SqlDbTypeVarBinary
	SqlDbTypeUniqueIdentifier
)

// SqlDbTypeFromName returns the SqlDbType for the given name.
func SqlDbTypeFromName(name string) SqlDbType {
	switch name {
	case "Bit":
		return SqlDbTypeBit
	case "TinyInt":
		return SqlDbTypeTinyInt
	case "Float":
		return SqlDbTypeFloat
	case "Double":
		return SqlDbTypeDouble
	case "Decimal":
		return SqlDbTypeDecimal
	case "Int":
		return SqlDbTypeInt
	case "BigInt":
		return SqlDbTypeBigInt
	case "DateTime":
		return SqlDbTypeDateTime
	case "DateTimeOffset":
		return SqlDbTypeDateTimeOffset
	case "Time":
		return SqlDbTypeTime
	case "VarChar":
		return SqlDbTypeVarChar
	case "NVarChar":
		return SqlDbTypeNVarChar
	case "Binary":
		return SqlDbTypeBinary
	case "VarBinary":
		return SqlDbTypeVarBinary
	case "UniqueIdentifier":
		return SqlDbTypeUniqueIdentifier
	default:
		return SqlDbTypeNVarChar
	}
}

// Name returns the name of the SqlDbType.
func (t SqlDbType) Name() string {
	switch t {
	case SqlDbTypeBit:
		return "Bit"
	case SqlDbTypeTinyInt:
		return "TinyInt"
	case SqlDbTypeFloat:
		return "Float"
	case SqlDbTypeDouble:
		return "Double"
	case SqlDbTypeDecimal:
		return "Decimal"
	case SqlDbTypeInt:
		return "Int"
	case SqlDbTypeBigInt:
		return "BigInt"
	case SqlDbTypeDateTime:
		return "DateTime"
	case SqlDbTypeDateTimeOffset:
		return "DateTimeOffset"
	case SqlDbTypeTime:
		return "Time"
	case SqlDbTypeVarChar:
		return "VarChar"
	case SqlDbTypeNVarChar:
		return "NVarChar"
	case SqlDbTypeBinary:
		return "Binary"
	case SqlDbTypeVarBinary:
		return "VarBinary"
	case SqlDbTypeUniqueIdentifier:
		return "UniqueIdentifier"
	default:
		return "NVarChar"
	}
}

// SqlDbTypeFuncMap is a map of SqlDbType to a function.
// Used in templates to generate code.
var SqlDbTypeFuncMap = template.FuncMap{
	"Bit": func() SqlDbType {
		return SqlDbTypeBit
	},
	"TinyInt": func() SqlDbType {
		return SqlDbTypeTinyInt
	},
	"Float": func() SqlDbType {
		return SqlDbTypeFloat
	},
	"Double": func() SqlDbType {
		return SqlDbTypeDouble
	},
	"Decimal": func() SqlDbType {
		return SqlDbTypeDecimal
	},
	"Int": func() SqlDbType {
		return SqlDbTypeInt
	},
	"BigInt": func() SqlDbType {
		return SqlDbTypeBigInt
	},
	"DateTime": func() SqlDbType {
		return SqlDbTypeDateTime
	},
	"DateTimeOffset": func() SqlDbType {
		return SqlDbTypeDateTimeOffset
	},
	"Time": func() SqlDbType {
		return SqlDbTypeTime
	},
	"VarChar": func() SqlDbType {
		return SqlDbTypeVarChar
	},
	"NVarChar": func() SqlDbType {
		return SqlDbTypeNVarChar
	},
	"Binary": func() SqlDbType {
		return SqlDbTypeBinary
	},
	"VarBinary": func() SqlDbType {
		return SqlDbTypeVarBinary
	},
	"UniqueIdentifier": func() SqlDbType {
		return SqlDbTypeUniqueIdentifier
	},
}

// SqlDbTypeFromCSharpType returns the SqlDbType for the given CSharpType.
func SqlDbTypeFromCSharpType(t CSharpType) SqlDbType {
	switch t {
	case CSharpTypeBool:
		return SqlDbTypeBit
	case CSharpTypeByte:
		return SqlDbTypeTinyInt
	case CSharpTypeFloat:
		return SqlDbTypeFloat
	case CSharpTypeDouble:
		return SqlDbTypeDouble
	case CSharpTypeDecimal:
		return SqlDbTypeDecimal
	case CSharpTypeInt:
		return SqlDbTypeInt
	case CSharpTypeLong:
		return SqlDbTypeBigInt
	case CSharpTypeDateTime:
		return SqlDbTypeDateTime
	case CSharpTypeDateTimeOffset:
		return SqlDbTypeDateTimeOffset
	case CSharpTypeString:
		return SqlDbTypeNVarChar
	case CSharpTypeByteArray:
		return SqlDbTypeBinary
	case CSharpTypeGuid:
		return SqlDbTypeUniqueIdentifier
	default:
		return SqlDbTypeNVarChar
	}
}

// MarshalYAML marshals the SqlDbType to YAML.
func (t SqlDbType) MarshalYAML() (interface{}, error) {
	return t.Name(), nil
}

// UnmarshalYAML unmarshals the SqlDbType from YAML.
func (t *SqlDbType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	if err := unmarshal(&name); err != nil {
		return err
	}
	*t = SqlDbTypeFromName(name)
	return nil
}

// MarshalJSON marshals the SqlDbType to JSON.
func (t SqlDbType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name())
}

// UnmarshalJSON unmarshals the SqlDbType from JSON.
func (t *SqlDbType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = SqlDbTypeFromName(name)
	return nil
}
