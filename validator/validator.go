package validator

type Validator struct {
}

type ValidateResult struct {
}

type ValidateParams struct {
}

func (v *Validator) Validate(params *ValidateParams) (*ValidateResult, error) {
  return &ValidateResult{}, nil
}
