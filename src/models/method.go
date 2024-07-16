package models

import "github.rbx.com/roblox/entity-schema-generator/enums"

// Method represents a method of an entity.
type Method struct {
	// Name of the method. This is required.
	Name string `json:"name" yaml:"name"`

	// Name of the dependent method for a must-get method. This is optional.
	DependentMethod string `json:"dependsOn" yaml:"depends_on"`

	// Name of the method in the DAL. If this is not specified, the method name
	// will default to the name of the method.
	DALName string `json:"dalName" yaml:"dal_name"`

	// Type of the method. This is required.
	Type enums.MethodType `json:"type" yaml:"type"`

	// Stored procedure name for this method. If this is not specified, the
	// stored procedure name will default to the name of the method.
	StoredProcedure string `json:"storedProcedure" yaml:"stored_procedure"`

	// The parameters of this method. This is required.
	Parameters []*Parameter `json:"parameters" yaml:"parameters"`

	// Additional parameters to pass to an ExclusiveStart method. This is optional
	// and only valid for GetCollectionExclusive methods.
	// This is for a legacy relic where some SQL methods could consolidate
	// multiple methods into one. This is not recommended for new methods.
	ExclusiveStartParameters []*Parameter `json:"exclusiveStartParameters" yaml:"exclusive_start_parameters"`

	// The passive properties of this method. This is optional.
	// These properties are used on the SQL side to denote these
	// properties as always being the same value in a WHERE clause.
	// e.g. GetValidAccountPasswordHashByAccountID without this only produces:
	// SELECT PasswordHash FROM Account WHERE AccountID = @AccountID
	// With this, it produces:
	// SELECT PasswordHash FROM Account WHERE AccountID = @AccountID AND @Valid = 1
	PassiveProperties []*PassiveProperty `json:"passiveProperties" yaml:"passive_properties"`

	// The visibility of this method. This is optional. Default is public.
	Visibility enums.VisibilityType `json:"visibility" yaml:"visibility"`

	// The type of the return value for count methods. This is only valid for
	// count methods. Default is int.
	CountReturnType enums.CSharpType `json:"countReturnType" yaml:"count_return_type"`

	// The constructed string for the method's parameters. This is not exposed
	// through the JSON or YAML.
	ConstructedStringParameters string `json:"-" yaml:"-"`

	// The name of the entity this method belongs to. This is not exposed through
	// the JSON or YAML.
	EntityName string `json:"-" yaml:"-"`

	// The name of the table this method belongs to. This is not exposed through
	// the JSON or YAML.
	Table string `json:"-" yaml:"-"`

	// The ID property of the entity this method belongs to. This is not exposed
	// through the JSON or YAML.
	IDProperty *Property `json:"-" yaml:"-"`

	// The lookup key for this method. This is not exposed through the JSON or
	// YAML.
	LookupKey string `json:"-" yaml:"-"`

	// The remote cacheable settings for this method. This is not exposed through
	// the JSON or YAML.
	RemoteCacheable *RemoteCacheable `json:"-" yaml:"-"`

	// The constructed list of parameters for the method. This is not exposed
	// through the JSON or YAML.
	ConstructedParameters string `json:"-" yaml:"-"`

	// The constructed list of parameters for the method call. This is not exposed
	// through the JSON or YAML.
	ConstructedParametersFormatted string `json:"-" yaml:"-"`

	// The constructed list of exclusive start parameters for the method. This is
	// not exposed through the JSON or YAML.
	ConstructedExclusiveStartStringParameters string `json:"-" yaml:"-"`

	// The constructed list of exclusive start parameters for the method call. This
	// is not exposed through the JSON or YAML.
	ConstructedExclusiveStartParameters string `json:"-" yaml:"-"`

	// The collection identifier for this method. This is not exposed through the
	// JSON or YAML.
	CollectionIdentifier string `json:"-" yaml:"-"`

	// The cache policy for this method. This is not exposed through the JSON or
	// YAML.
	CachePolicy string `json:"-" yaml:"-"`

	// The version of the entity. This is not exposed through the JSON or YAML.
	Version int `json:"-" yaml:"-"`

	// The properties of the entity. This is not exposed through the JSON or YAML.
	Properties []*Property `json:"-" yaml:"-"`

	// The type of entity. This is not exposed through the JSON or YAML.
	EntityType string `json:"-" yaml:"-"`
}
