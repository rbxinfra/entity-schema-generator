package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/models"
	"github.rbx.com/roblox/entity-schema-generator/templates/platform"
)

func parseTemplateForPlatformEntityInterface(model *models.Entity) (string, error) {
	tmpl := template.New("PlatformEntityInterface")
	applyCustomFunctions(tmpl)

	var err error
	if tmpl, err = tmpl.Parse(string(platform.InterfaceTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}

func parseTemplateForPlatformEntityImplementation(model *models.Entity) (string, error) {
	tmpl := template.New("PlatformEntityImplementation")
	applyCustomFunctions(tmpl)

	var err error
	if tmpl, err = tmpl.Parse(string(platform.ImplementationTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}
