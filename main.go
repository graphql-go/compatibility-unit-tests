package main

import (
	"log"

	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"
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
	result, err := app.Run(cliResult.Choice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("result: %+v", result)
}
