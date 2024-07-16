package templates

import (
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/enums"
)

func appendEnums(funcMap template.FuncMap) template.FuncMap {
	for key, f := range enums.MethodTypeFuncMap {
		funcMap[key] = f
	}

	for key, f := range enums.CSharpTypeFuncMap {
		funcMap[key] = f
	}

	for key, f := range enums.SqlDbTypeFuncMap {
		funcMap[key] = f
	}

	return funcMap
}
