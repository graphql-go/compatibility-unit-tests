package main

import (
	"fmt"
	"log"

	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"
)

var choices = []string{}

func init() {
	for _, i := range implementation.Implementations {
		choices = append(choices, i.Repo.URL)
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

	currentImplementation, ok := implementation.ImplementationsMap[cliResult.Choice]
	if !ok {
		log.Fatal(fmt.Errorf("expected to find the following implementation: %v", cliResult.Choice))
	}

	app := mainApp.App{}
	result, err := app.Run(mainApp.AppParams{
		Implementation:    currentImplementation,
		RefImplementation: implementation.GraphqlJSImplementation,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("result: %+v", result)
}
