package templates

import (
	"strings"
	"text/template"

	"github.rbx.com/roblox/entity-schema-generator/models"
	"github.rbx.com/roblox/entity-schema-generator/templates/mssql"
)

// ParseTemplateForDatabase parses a template for a given database.
func ParseTemplateForDatabase(model *models.Database) (string, error) {
	tmpl := template.New("Database")
	applyCustomFunctions(tmpl)
	var err error
	if tmpl, err = tmpl.Parse(string(mssql.DatabaseTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.LookupTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.GetOrCreateTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.GetCollectionTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.GetCollectionPagedTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.GetCountTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.MultiGetTemplate)); err != nil {
		return "", err
	}
	if tmpl, err = tmpl.Parse(string(mssql.GetCollectionExclusiveTemplate)); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err := tmpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return postProcess(textWriter.String()), nil
}
