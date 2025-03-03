package validator

import (
	"graphql-go/compatibility-unit-tests/types"
)

type Validator struct {
}

type ValidatorParams struct {
	ImplementationTests    []types.ImplementationTest
	RefImplementationTests []types.ImplementationTest
}

type ValidatorResult struct {
	SuccessfulTests []types.SuccessfulTest
	FailedTests     []types.FailedTest
}

func (v *Validator) Validate(params *ValidatorParams) (*ValidatorResult, error) {
	return &ValidatorResult{
		SuccessfulTests: []types.SuccessfulTest{},
		FailedTests:     []types.FailedTest{},
	}, nil
}
