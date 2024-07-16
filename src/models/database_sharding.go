package models

// DatabaseSharding represents a sharded database.
type DatabaseSharding struct {
	// ShardCount of the shards. Default is 8.
	ShardCount *int `json:"count" yaml:"count"`

	// Is the sharding server on IsWindows. Default is false.
	IsWindows *bool `json:"windows" yaml:"windows"`

	// The initial size of the shard as a size string. Default is 5120KB.
	InitialSize string `json:"initialSize" yaml:"initial_size"`

	// The maximum size of the shard as a size string. Default is 10MB.
	MaximumSize string `json:"maximumSize" yaml:"maximum_size"`

	// The growth increment of the shard as a size string. Default is 1024KB.
	GrowthSize string `json:"growthSize" yaml:"growth_size"`
}
