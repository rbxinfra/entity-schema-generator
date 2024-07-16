package mssql

import _ "embed"

//go:embed dal.txt
var DalTemplate []byte

//go:embed get_or_create.txt
var GetOrCreateTemplate []byte

//go:embed lookup.txt
var LookupTemplate []byte

//go:embed get_collection.txt
var GetCollectionTemplate []byte

//go:embed get_collection_paged.txt
var GetCollectionPagedTemplate []byte

//go:embed get_count.txt
var GetCountTemplate []byte

//go:embed multi_get.txt
var MultiGetTemplate []byte

//go:embed get_collection_exclusive.txt
var GetCollectionExclusiveTemplate []byte
