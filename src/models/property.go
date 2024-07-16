package models

import "github.rbx.com/roblox/entity-schema-generator/enums"

// Property represents a property of an entity.
type Property struct {
	// Name of the property. This is required.
	Name string `json:"name" yaml:"name"`

	// The description of the property. This is optional.
	// Only used for the platform layer.
	Description string `json:"description" yaml:"description"`

	// A key this key points to. This is optional.
	// Required format: [{{database}}?].[dbo].[{{table}}].[{{column}}]
	// If database is not specified, it will default to the current database.
	ForeignKey string `json:"foreignKey" yaml:"foreign_key"`

	// Type of the property for C#. This is required.
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

	// Is this property read-only? This is optional. Default is true.
	ReadOnly *bool `json:"readOnly" yaml:"read_only"`

	// Is this property nullable? This is optional. Default is true.
	Nullable *bool `json:"nullable" yaml:"nullable"`

	// The visibility of this property. This is optional. Default is public.
	Visibility enums.VisibilityType `json:"visibility" yaml:"visibility"`

	// The SqlDbType of this property.
	SqlDbType enums.SqlDbType `json:"sqlDbType" yaml:"sql_db_type"`

	// The constructed string for the foreign key constraint. This is not exposed through the JSON or YAML.
	ConstructedForeignKeyConstraintKey string `json:"-" yaml:"-"`

	// The constructed string for the foreign key. This is not exposed through the JSON or YAML.
	// This is constructed here because it has some logic to determine if the foreign key is outside the current database.
	ConstructedForeignKey string `json:"-" yaml:"-"`
}
