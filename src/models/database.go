package models

import "github.rbx.com/roblox/entity-schema-generator/enums"

// Database represents a database.
type Database struct {
	// Name of the database. This is required.
	Name string `json:"name" yaml:"name"`

	// The sharding configuration for this database. If nil, sharding is disabled.
	DatabaseSharding *DatabaseSharding `json:"sharding" yaml:"sharding"`

	// The data tables in this database. This is not exposed through the JSON or YAML.
	// Based on the type of the ID property, and if a multi-get exists.
	DataTables []enums.SqlDbType `json:"-" yaml:"-"`

	// The entities in this database. This is not exposed through the JSON or YAML.
	Entities []*Entity `json:"-" yaml:"-"`
}
