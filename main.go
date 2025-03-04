package main

import (
	"fmt"
	"log"

	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"
	"graphql-go/compatibility-unit-tests/result"
)

var choices = []string{}

func init() {
	for _, i := range implementation.Implementations {
		choices = append(choices, i.Repo.String(implementation.ImplementationPrefix))
	}
}

func main() {
	cli := cmd.CLI{}

	header := implementation.RefImplementation.Repo.String(implementation.RefImplementationPrefix)

	cliResult, err := cli.Run(&cmd.RunParams{
		Choices: choices,
		Header:  header,
	})
	if err != nil {
		log.Fatal(err)
	}

	currentImplementation, ok := implementation.ImplementationsMap[cliResult.Choice]
	if !ok {
		log.Fatal(fmt.Errorf("expected to find the following implementation: %v", cliResult.Choice))
	}

	app := mainApp.App{}
	r, err := app.Run(mainApp.AppParams{
		Implementation:    currentImplementation,
		RefImplementation: implementation.GraphqlJSImplementation,
	})
	if err != nil {
		log.Fatal(err)
	}

	summaryParams := &result.SummaryParams{
		SuccessfulTests: len(r.SuccessfulTests),
		FailedTests:     len(r.FailedTests),
	}
	result := result.Result{}

	fmt.Println(result.Summary(summaryParams))
}
