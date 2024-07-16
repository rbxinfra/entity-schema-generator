package mssql

import _ "embed"

//go:embed database.txt
var DatabaseTemplate string

//go:embed lookup.txt
var LookupTemplate string

//go:embed get_or_create.txt
var GetOrCreateTemplate string

//go:embed get_collection.txt
var GetCollectionTemplate string

//go:embed get_collection_paged.txt
var GetCollectionPagedTemplate string

//go:embed get_count.txt
var GetCountTemplate string

//go:embed multi_get.txt
var MultiGetTemplate string

//go:embed get_collection_exclusive.txt
var GetCollectionExclusiveTemplate string
