package app

import (
	"graphql-go/compatibility-unit-tests/extractor"
	"graphql-go/compatibility-unit-tests/puller"
	"graphql-go/compatibility-unit-tests/types"
	"graphql-go/compatibility-unit-tests/validator"
)

type App struct {
}

type AppResult struct {
	SuccessfulTests []types.SuccessfulTest
	FailedTests     []types.FailedTest
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
		ImplementationType:    params.Implementation.Type,
		RefImplementationType: params.RefImplementation.Type,
		RootDir:               params.RefImplementation.Repo.Dir,
	}); err != nil {
		return nil, err
	}

	val := validator.Validator{}
	validatorResult, err := val.Validate(&validator.ValidatorParams{
		ImplementationTests:    []types.ImplementationTest{},
		RefImplementationTests: []types.ImplementationTest{},
	})
	if err != nil {
		return nil, err
	}

	return &AppResult{
		SuccessfulTests: validatorResult.SuccessfulTests,
		FailedTests:     validatorResult.FailedTests,
	}, nil
}
