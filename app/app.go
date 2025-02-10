package app

import (
	"graphql-go/compatibility-unit-tests/puller"
)

type App struct {
}

type AppResult struct {
}

func (app *App) Run(repoURL string) (*AppResult, error) {
	p := puller.Puller{}

	if _, err := p.Pull(&puller.PullerParams{
		RepoURL: repoURL,
	}); err != nil {
		return nil, err
	}

	return &AppResult{}, nil
}
