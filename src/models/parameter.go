package models

import "github.rbx.com/roblox/entity-schema-generator/enums"

// Parameter represents a parameter of an entity method.
type Parameter struct {
	// Name of the parameter. This is required.
	Name string `json:"name" yaml:"name"`

	// Type of the parameter for C#. This is required.
	Type enums.CSharpType `json:"type" yaml:"type"`

	// The length of a string property. This is optional.
	// If not specified or set to 0, the length will be MAX.
	Length int `json:"length" yaml:"length"`

	// Is the binary property a varbinary? This is optional. Default is false.
	IsVarBinary *bool `json:"isVarBinary" yaml:"is_var_binary"`

	// Is the string property an unicode string? This is optional. Default is true.
	IsUnicode *bool `json:"isUnicode" yaml:"is_unicode"`

	// Is the DateTime property in UTC? This is optional. Default is false.
	IsUTC *bool `json:"isUTC" yaml:"is_utc"`

	// The SqlDbType of the parameter for SQL Server. This is not exposed through YAML or JSON.
	SqlDbType enums.SqlDbType `json:"-" yaml:"-"`
}
