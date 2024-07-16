package main

import (
	"flag"
	"os"
	"path"

	"github.rbx.com/roblox/entity-schema-generator/configuration"
	"github.rbx.com/roblox/entity-schema-generator/flags"
	"github.rbx.com/roblox/entity-schema-generator/templates"
)

var applicationName string
var buildMode string
var commitSha string

// Pre-setup, runs before main.
func init() {
	flags.SetupFlags(applicationName, buildMode, commitSha)
}

// Main entrypoint.
func main() {
	if len(os.Args) == 1 {
		flag.Usage()

		return
	}

	if *flags.HelpFlag {
		flag.Usage()

		return
	}

	entities, databases, err := configuration.Parse()
	if err != nil {
		panic(err)
	}

	for _, entity := range entities {
		var bllText, dalText, platformInterface, platformImplementation string
		if bllText, dalText, platformInterface, platformImplementation, err = templates.ParseTemplateForEntity(entity); err != nil {
			panic(err)
		}

		pathName := path.Join(*flags.OutputDirectoryFlag, entity.Database, entity.EntityName)
		if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
			panic(err)
		}

		bllFileName := path.Join(pathName, entity.EntityName+".cs")
		dalFileName := path.Join(pathName, entity.EntityName+"DAL.cs")
		platformInterfaceFileName := path.Join(pathName, "I"+entity.EntityName+"Entity.cs")
		platformImplementationFileName := path.Join(pathName, entity.EntityName+"Entity.cs")

		if err := os.WriteFile(bllFileName, []byte(bllText), os.ModePerm); err != nil {
			panic(err)
		}

		if err := os.WriteFile(dalFileName, []byte(dalText), os.ModePerm); err != nil {
			panic(err)
		}

		if err := os.WriteFile(platformInterfaceFileName, []byte(platformInterface), os.ModePerm); err != nil {
			panic(err)
		}

		if err := os.WriteFile(platformImplementationFileName, []byte(platformImplementation), os.ModePerm); err != nil {
			panic(err)
		}

		println("Generated files for entity", entity.EntityName)
	}

	for _, database := range databases {
		var databaseText string
		if databaseText, err = templates.ParseTemplateForDatabase(database); err != nil {
			panic(err)
		}

		pathName := *flags.OutputDirectoryFlag
		if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
			panic(err)
		}

		databaseFileName := path.Join(pathName, database.Name+".sql")

		if err := os.WriteFile(databaseFileName, []byte(databaseText), os.ModePerm); err != nil {
			panic(err)
		}

		println("Generated files for database", database.Name)
	}
}
