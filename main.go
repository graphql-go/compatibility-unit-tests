package main

import (
	"log"

	"graphql-go/compatibility-unit-tests/cmd"
	"graphql-go/compatibility-unit-tests/implementation"
	"graphql-go/compatibility-unit-tests/puller"
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
		Choices: implementation.Implementations,
	})
	if err != nil {
		log.Fatal(err)
	}

	p := puller.Puller{}
	result, err := p.Pull(&puller.PullerParams{
		RepoURL: cliResult.Choice,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("result: %+v", result)
}
