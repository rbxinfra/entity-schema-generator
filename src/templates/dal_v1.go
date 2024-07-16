package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/models"
	dataaccessv1 "github.rbx.com/roblox/entity-schema-generator/templates/dal/data_access_v1"
)

func parseTemplateForEntityDalV1(model *models.Entity) (string, error) {
	tmpl := template.New("DALv1")
	applyCustomFunctions(tmpl)

	var err error
	if tmpl, err = tmpl.Parse(string(dataaccessv1.DalTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.GetOrCreateTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.LookupTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.GetCollectionTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.GetCollectionPagedTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.GetCountTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.MultiGetTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessv1.GetCollectionExclusiveTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}
