package v1

import _ "embed"

//go:embed bll.txt
var BllTemplate []byte

//go:embed lookup.txt
var LookupTemplate []byte

//go:embed get_collection_paged.txt
var GetCollectionPagedTemplate []byte

//go:embed get_collection.txt
var GetCollectionTemplate []byte

//go:embed multi_get.txt
var MultiGetTemplate []byte

//go:embed get_count.txt
var GetCountTemplate []byte

//go:embed get_or_create.txt
var GetOrCreateTemplate []byte

//go:embed must_get.txt
var MustGetTemplate []byte

//go:embed get_collection_exclusive.txt
var GetCollectionExclusiveTemplate []byte
