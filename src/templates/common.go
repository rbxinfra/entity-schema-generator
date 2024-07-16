package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/helpers"
)

func applyCustomFunctions(tmpl *template.Template) *template.Template {
	funcMap := template.FuncMap{
		"toCamelCase":                 helpers.ToCamelCase,
		"toPascalCase":                helpers.ToPascalCase,
		"toJson":                      helpers.ToJson,
		"toLower":                     strings.ToLower,
		"deref":                       func(b *bool) bool { return *b },
		"derefInt":                    func(i *int) int { return *i },
		"toStr":                       helpers.ToString,
		"getRequiredWhitespaceMarker": getRequiredWhitespaceMarker,
		"split":                       strings.Split,
		"getFirstElement":             helpers.GetFirstElement,
		"normalizePascalParts":        helpers.NormalizePascalParts,
		"loop":                        helpers.Loop,
		"sub":                         func(i, j int) int { return i - j },
	}

	funcMap = appendEnums(funcMap)

	return tmpl.Funcs(funcMap)
}
