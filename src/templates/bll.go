package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/models"
	bllv1 "github.rbx.com/roblox/entity-schema-generator/templates/bll/bll_v1"
)

func parseTemplateForEntityBllV1(model *models.Entity) (string, error) {
	tmpl := template.New("BLLv1")
	applyCustomFunctions(tmpl)

	var err error
	if tmpl, err = tmpl.Parse(string(bllv1.BllTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.LookupTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.GetCollectionPagedTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.GetCollectionTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.MultiGetTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.GetCountTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.GetOrCreateTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.MustGetTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(bllv1.GetCollectionExclusiveTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}
