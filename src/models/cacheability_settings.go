package models

import "fmt"

// CacheabilitySettings represents the cacheability settings
// for an entity.
type CacheabilitySettings struct {
	// Are collections of this entity cacheable?
	CollectionsAreCacheable bool `json:"collectionsAreCacheable" yaml:"collections_are_cacheable"`

	// Are counts of this entity cacheable?
	CountsAreCacheable bool `json:"countsAreCacheable" yaml:"counts_are_cacheable"`

	// Is this entity cacheable?
	EntityIsCacheable bool `json:"entityIsCacheable" yaml:"entity_is_cacheable"`

	// Are id lookups of this entity cacheable?
	IDLookupsAreCacheable bool `json:"idLookupsAreCacheable" yaml:"id_lookups_are_cacheable"`

	// Does this entity have unqualified collections?
	HasUnqualifiedCollections bool `json:"hasUnqualifiedCollections" yaml:"has_unqualified_collections"`

	// Are id lookups of this entity case sensitive?
	IDLookupsAreCaseSensitive bool `json:"idLookupsAreCaseSensitive" yaml:"id_lookups_are_case_sensitive"`

	// Is null cacheable?
	IsNullCacheable bool `json:"isNullCacheable" yaml:"is_null_cacheable"`
}

// DefaultCacheabilitySettings returns the default cacheability settings.
func DefaultCacheabilitySettings() *CacheabilitySettings {
	return &CacheabilitySettings{
		CollectionsAreCacheable:   false,
		CountsAreCacheable:        false,
		EntityIsCacheable:         true,
		IDLookupsAreCacheable:     true,
		HasUnqualifiedCollections: false,
		IDLookupsAreCaseSensitive: false,
		IsNullCacheable:           true,
	}
}

// String returns the string representation of the cacheability settings.
func (s *CacheabilitySettings) String() string {
	return fmt.Sprintf(
		"collectionsAreCacheable: %t, countsAreCacheable: %t, entityIsCacheable: %t, idLookupsAreCacheable: %t, hasUnqualifiedCollections: %t, idLookupsAreCaseSensitive: %t",
		s.CollectionsAreCacheable,
		s.CountsAreCacheable,
		s.EntityIsCacheable,
		s.IDLookupsAreCacheable,
		s.HasUnqualifiedCollections,
		s.IDLookupsAreCaseSensitive,
	)
}
