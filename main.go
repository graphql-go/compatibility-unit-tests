package main

import (
	"fmt"
	"log"
	"strconv"

	mainApp "graphql-go/compatibility-unit-tests/app"
	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"

	"github.com/charmbracelet/lipgloss"
)

var choices = []string{}

var successfulStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#4CAF50")).
	PaddingTop(0).
	PaddingLeft(0).
	Width(40)

var failedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#A52A2A")).
	PaddingTop(0).
	PaddingLeft(0).
	Width(40)

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

	sTests := strconv.Itoa(len(result.SuccessfulTests))
	fTests := strconv.Itoa(len(result.FailedTests))

	fmt.Println(successfulStyle.Render(fmt.Sprintf("successful tests count: %+v", sTests)))
	fmt.Println(failedStyle.Render(fmt.Sprintf("failed tests count: %+v", fTests)))
}
