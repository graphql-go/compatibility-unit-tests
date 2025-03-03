package validator

import (
	"graphql-go/compatibility-unit-tests/types"
)

type Validator struct {
}

type ValidatorParams struct {
	ImplementationTests    types.Implementation
	RefImplementationTests types.Implementation
}

type ValidatorResult struct {
	SuccessfulTests []types.SuccessfulTest
	FailedTests     []types.FailedTest
}

func (v *Validator) Validate(params *ValidatorParams) (*ValidatorResult, error) {
	refImplTestsMap := make(map[string]string, 0)

	successfultTests := []types.SuccessfulTest{}
	failedTests := []types.FailedTest{}

	for _, testName := range params.RefImplementationTests.TestNames {
		refImplTestsMap[testName] = testName
	}

	for _, testName := range params.ImplementationTests.TestNames {
		tName, found := refImplTestsMap[testName]
		if found {
			successfultTests = append(successfultTests, types.SuccessfulTest{
				Name: tName,
			})
		} else {
			failedTests = append(failedTests, types.FailedTest{
				Name: tName,
			})
		}
	}

	return &ValidatorResult{
		SuccessfulTests: successfultTests,
		FailedTests:     failedTests,
	}, nil
}
