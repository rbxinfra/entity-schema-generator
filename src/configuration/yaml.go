package configuration

import (
	"os"

	"github.rbx.com/roblox/entity-schema-generator/models"
	"gopkg.in/yaml.v2"
)

func parseEntityYAMLFile(fileName string) (*models.Entity, error) {
	var entity models.Entity

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	yamlParser := yaml.NewDecoder(yamlFile)
	if err = yamlParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseDatabaseYAMLFile(fileName string) (*models.Database, error) {
	var database models.Database

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	yamlParser := yaml.NewDecoder(yamlFile)
	if err = yamlParser.Decode(&database); err != nil {
		return nil, err
	}

	return &database, nil
}
