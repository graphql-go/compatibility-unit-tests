package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/compatibility-base/bubbletea"
	"github.com/graphql-go/compatibility-base/cmd"
	"github.com/graphql-go/compatibility-base/config"
	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/implementation"
	"graphql-go/compatibility-unit-tests/result"
)

func main() {
	cfg := config.New()

	choicesModelUIHeader := cfg.GraphqlJSImplementation.Repo.String(implementation.RefImplementationPrefix)

	cmdParams := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{
			Models: bubbletea.Models{
				bubbletea.NewChoicesModel(&bubbletea.ChoicesModelParams{
					Order:   1,
					Choices: cfg.AvailableImplementations,
					UI: bubbletea.ChoicesModelUIParams{
						Header: choicesModelUIHeader,
					},
				}),
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}
	cli := cmd.New(&cmdParams)

	runParams := &cmd.RunParams{
		ResultCallback: nil,
	}

	runResult, err := cli.Run(runParams)
	if err != nil {
		log.Printf("failed to run cli: %v", err)
	}

	currentImplementation, ok := implementation.ImplementationsMap[runResult.ChoicesModelResult.Choice]
	if !ok {
		log.Fatal(fmt.Errorf("expected to find the following implementation: %v", runResult.ChoicesModelResult.Choice))
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
