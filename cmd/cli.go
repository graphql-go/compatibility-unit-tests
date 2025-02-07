package cmd

import (
	"graphql-go/compatibility-unit-tests/bubbletea"
)

type CLI struct {
}

type model struct {
}

type RunResult struct {
	Choice string
}

type RunParams struct {
	Choices []string
}

func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	bt := bubbletea.NewBubbleTea(&bubbletea.BubbleTeaParams{
		Choices: p.Choices,
	})

	btRunResult, err := bt.Run()
	if err != nil {
		return nil, err
	}

	return &RunResult{
		Choice: btRunResult.Choice,
	}, nil
}
