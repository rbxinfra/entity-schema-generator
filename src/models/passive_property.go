package models

// PassiveProperty represents a passive property of an entity.
type PassiveProperty struct {
	// Name of the property. This is required.
	// It must match one of the properties in the entity.
	// It can not be the ID property.
	// It can also not be the name of an argument on the method.
	Name string `json:"name" yaml:"name"`

	// The passive value of the property. This is required.
	// The SQL value of the property is parsed into this value.
	Value string `json:"value" yaml:"value"`
}
