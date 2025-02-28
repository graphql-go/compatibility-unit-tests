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
	extractResult, err := ex.Extract(&extractor.ExtractorParams{
		RootDir: params.RefImplementation.Repo.Dir,
	})
	if err != nil {
		return nil, err
	}

	val := validator.Validator{}
	validatorResult, err := val.Validate(&validator.ValidatorParams{
		ImplementationTests:    extractResult.TestNames,
		RefImplementationTests: extractResult.TestNames,
	})
	if err != nil {
		return nil, err
	}

	return &AppResult{}, nil
}
