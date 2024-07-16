package configuration

import (
	"encoding/json"
	"os"

	"github.rbx.com/roblox/entity-schema-generator/models"
)

func parseEntityJSONFile(fileName string) (*models.Entity, error) {
	var entity models.Entity

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseDatabaseJSONFile(fileName string) (*models.Database, error) {
	var database models.Database

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&database); err != nil {
		return nil, err
	}

	return &database, nil
}
