package models

// RemoteCacheable represents information about whether a model is in a remote cache.
type RemoteCacheable struct {
	// The name of the MemcacheD group setting. This is optional.
	// In the form: <ClassName>:<PropertyName>
	MemcachedGroupSetting string `json:"memcachedGroupSetting" yaml:"memcached_group_setting"`

	// The name of the settings class that contains the MemcacheD group setting. This is not exposed through the JSON or YAML.
	MemcachedGroupSettingClass string `json:"-" yaml:"-"`

	// The name of the property that contains the MemcacheD group setting. This is not exposed through the JSON or YAML.
	MemcachedGroupSettingProperty string `json:"-" yaml:"-"`
}
