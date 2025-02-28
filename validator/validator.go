package validator

type Validator struct {
}

type ValidateResult struct {
}

type ValidateParams struct {
	ImplementationTests    []types.Tests
	RefImplementationTests []types.Tests
}

func (v *Validator) Validate(params *ValidateParams) (*ValidateResult, error) {
	return &ValidateResult{}, nil
}
