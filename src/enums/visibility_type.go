package enums

import "encoding/json"

// VisibilityType is the visibility of a property.
type VisibilityType uint8

const (
	// Unknown is the default visibility.
	VisibilityTypeUnknown VisibilityType = iota

	// Public is the default visibility.
	VisibilityTypePublic

	// Internal can only be accessed from within the same assembly.
	VisibilityTypeInternal

	// Protected can only be accessed from a derived class.
	VisibilityTypeProtected
)

// VisibilityFromName returns the Visibility for the given name.
func VisibilityFromName(name string) VisibilityType {
	switch name {
	case "public":
		return VisibilityTypePublic
	case "internal":
		return VisibilityTypeInternal
	case "protected":
		return VisibilityTypeProtected
	default:
		return VisibilityTypeUnknown
	}
}

// Name returns the name of the Visibility.
func (v VisibilityType) Name() string {
	switch v {
	case VisibilityTypePublic:
		return "public"
	case VisibilityTypeInternal:
		return "internal"
	case VisibilityTypeProtected:
		return "protected"
	default:
		return "unknown"
	}
}

// MarshalYAML marshals the Visibility to YAML.
func (v VisibilityType) MarshalYAML() (interface{}, error) {
	return v.Name(), nil
}

// UnmarshalYAML unmarshals the Visibility from YAML.
func (v *VisibilityType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	if err := unmarshal(&name); err != nil {
		return err
	}

	*v = VisibilityFromName(name)
	return nil
}

// MarshalJSON marshals the Visibility to JSON.
func (v VisibilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Name())
}

// UnmarshalJSON unmarshals the Visibility from JSON.
func (v *VisibilityType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	*v = VisibilityFromName(name)
	return nil
}
