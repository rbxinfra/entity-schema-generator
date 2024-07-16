package models

// PredefinedCondition represents a predefined condition.
type PredefinedCondition struct {
	// The name of the condition. This is required.
	Name string `json:"name" yaml:"name"`

	// The props of the condition. This is optional.
	// The reason it is a map, even though only the values are used, is because
	// the keys are used in the schema to create clarity.
	// If none specified, it will use the name as the argument.
	Properties []map[string]string `json:"properties" yaml:"properties"`

	// The constructed arguments for the call to the get method. This is not exposed through the JSON or YAML.
	// Constructed based on the values.
	PropertiesConstructed string `json:"-" yaml:"-"`
}

// Predefined represents a predefined entity.
// These are entities that are predefined in the system.
// This system just determines how to generate the code for these entities.
type Predefined struct {
	// The name of the method to call to get the entity. This is required.
	GetMethod string `json:"method" yaml:"method"`

	// The conditions of the entity. This is required.
	Values []*PredefinedCondition `json:"values" yaml:"values"`
}
