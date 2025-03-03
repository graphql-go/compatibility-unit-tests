package app

import (
	"fmt"

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
	extractorResult, err := ex.Extract(&extractor.ExtractorParams{
		Implementation:    params.Implementation,
		RefImplementation: params.RefImplementation,
	})
	if err != nil {
		return nil, err
	}

	implementationTests, ok := extractorResult.TestNames[types.GoImplementationType]
	if !ok {
		return nil, fmt.Errorf("failed to find implementation type with key: %v", types.GoImplementationType)
	}

	refImplementationTests, ok := extractorResult.TestNames[types.RefImplementationType]
	if !ok {
		return nil, fmt.Errorf("failed to find implementation type with key: %v", types.RefImplementationType)
	}

	val := validator.Validator{}
	validatorResult, err := val.Validate(&validator.ValidatorParams{
		ImplementationTests:    implementationTests,
		RefImplementationTests: refImplementationTests,
	})
	if err != nil {
		return nil, err
	}

	return &AppResult{
		SuccessfulTests: validatorResult.SuccessfulTests,
		FailedTests:     validatorResult.FailedTests,
	}, nil
}
