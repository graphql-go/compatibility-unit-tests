package app

import (
	"graphql-go/compatibility-unit-tests/extractor"
	"graphql-go/compatibility-unit-tests/puller"
	"graphql-go/compatibility-unit-tests/types"
)

type App struct {
}

type AppResult struct {
}

type AppParams struct {
	Implementation    types.Implementation
	RefImplementation types.Implementation
}

func (app *App) Run(params AppParams) (*AppResult, error) {
	p := puller.Puller{}

	if _, err := p.Pull(&puller.PullerParams{
		Implementation:    params.Implementation,
		RefImplementation: params.RefImplementation,
	}); err != nil {
		return nil, err
	}

	ex := extractor.Extractor{}
	if _, err := ex.Extract(&extractor.ExtractorParams{
		Source: "",
	}); err != nil {
		return nil, err
	}

	return &AppResult{}, nil
}
