package main

import (
	"log"

	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"
	"graphql-go/compatibility-unit-tests/types"
)

var choices = []string{}

func init() {
	for i := range implementation.Implementations {
		choices = append(choices, i)
	}
}

func main() {
	cli := cmd.CLI{}

	cliResult, err := cli.Run(&cmd.RunParams{
		Choices: choices,
	})
	if err != nil {
		log.Fatal(err)
	}

	app := mainApp.App{}
	result, err := app.Run(mainApp.AppParams{
		Implementation: types.Implementation{
			Repo: types.Repository{
				Name: "implementation",
				URL:  cliResult.Choice,
			},
		},
		RefImplementation: types.Implementation{
			Repo: types.Repository{
				Name: "reference-implementation",
				URL:  implementation.JSRefImplementation,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("result: %+v", result)
}
