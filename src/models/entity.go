package models

import "github.rbx.com/roblox/entity-schema-generator/enums"

// Entity represents an entity.
type Entity struct {
	// Name of the entity. This is required.
	EntityName string `json:"name" yaml:"name"`

	// Name of the database this entity belongs to. This is required.
	Database string `json:"database" yaml:"database"`

	// Name of the table this entity is part of. This is required.
	Table string `json:"table" yaml:"table"`

	// The predefined values for this entity. This is optional.
	Predefined *Predefined `json:"predefined" yaml:"predefined"`

	// Name of the C# namespace this entity belongs to. If not specified, the
	// namespace will default to "Roblox".
	Namespace string `json:"namespace" yaml:"namespace"`

	// Name of the C# namespace the DAL for this entity belongs to. If not
	// specified, it will remain in the same namespace as the entity.
	DALNamespace string `json:"dalNamespace" yaml:"dal_namespace"`

	// Version of the entity. Defaults to 2 (MssqlDatabase).
	Version *int `json:"version" yaml:"version"`

	// The ID property of this entity. This is required.
	IDProperty *Property `json:"id" yaml:"id"`

	// Determines if a must-get method should be generated for this entity. This
	// is optional. Default is false.
	GenerateMustGet *bool `json:"generateMustGet" yaml:"generate_must_get"`

	// Determines if a create-new method should be generated for this entity. This
	// is optional. Default is false.
	GenerateCreateNew *bool `json:"generateCreateNew" yaml:"generate_create_new"`

	// The properties of this entity. This is required.
	Properties []*Property `json:"properties" yaml:"properties"`

	// The properties as an argument list. This is not exposed through the JSON or YAML.
	// Constructed based on the properties of the entity.
	PropertiesArgs string `json:"-" yaml:"-"`

	// The methods of this entity. This is required.
	Methods []*Method `json:"methods" yaml:"methods"`

	// The cacheability settings for this entity. This is optional. Default is
	// CacheabilityUnknown.
	CacheabilitySettings *CacheabilitySettings `json:"cacheabilitySettings" yaml:"cacheability_settings"`

	// The remote cacheability settings for this entity. This is optional.
	RemoteCacheable *RemoteCacheable `json:"remoteCacheable" yaml:"remote_cacheable"`

	// The visibility of this entity. This is optional. Default is public.
	Visibility enums.VisibilityType `json:"visibility" yaml:"visibility"`

	// The visibility of this entity's DAL. This is optional. Default is public.
	DALVisibility enums.VisibilityType `json:"dalVisibility" yaml:"dal_visibility"`

	// The list of lookup keys for this entity. This is not exposed through the JSON or YAML.
	// Constructed based on Lookup methods.
	LookupKeys []string `json:"-" yaml:"-"`

	// The list of state tokens for this entity. This is not exposed through the JSON or YAML.
	// Constructed based on State methods such as collections and counts.
	StateTokens []string `json:"-" yaml:"-"`

	// A dictionary used to prevent duplicate lookup keys. This is not exposed through the JSON or YAML.
	// Constructed based on Lookup methods.
	LookupKeyMap map[string]bool `json:"-" yaml:"-"`

	// A dictionary used to prevent duplicate state tokens. This is not exposed through the JSON or YAML.
	// Constructed based on State methods such as collections and counts.
	StateTokenMap map[string]bool `json:"-" yaml:"-"`

	// The type of the entity. This is not exposed through the JSON or YAML.
	// Constructed based on the version of the entity.
	EntityType string `json:"-" yaml:"-"`

	// Is the updated property nullable. This is not exposed through the JSON or YAML.
	IsUpdatedNullable bool `json:"-" yaml:"-"`
}
