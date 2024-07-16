package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/models"
	dataaccessmssql "github.rbx.com/roblox/entity-schema-generator/templates/dal/data_access_mssql"
)

func parseTemplateForEntityDalV2(model *models.Entity) (string, error) {
	tmpl := template.New("DALv2")
	applyCustomFunctions(tmpl)

	var err error
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.DalTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.GetOrCreateTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.LookupTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.GetCollectionTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.GetCollectionPagedTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.GetCountTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.MultiGetTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(dataaccessmssql.GetCollectionExclusiveTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}
