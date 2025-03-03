package extractor

import (
	"fmt"

	"graphql-go/compatibility-unit-tests/types"
)

type Extractor struct {
}

type ExtractorParams struct {
	Implementation    types.Implementation
	RefImplementation types.Implementation
}

type ExtractorResult struct {
	TestNames map[types.ImplementationType]types.Implementation
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	testNames, err := e.implementationTestNames(ImplTestNamesParams{
		Implementations: []types.Implementation{params.Implementation, params.RefImplementation},
	})
	if err != nil {
		return nil, err
	}

	return &ExtractorResult{
		TestNames: testNames,
	}, nil
}

type ImplTestNamesParams struct {
	Implementations []types.Implementation
}

func (e *Extractor) implementationTestNames(params ImplTestNamesParams) (map[types.ImplementationType]types.Implementation, error) {
	result := map[types.ImplementationType]types.Implementation{}

	for i := 0; i < len(params.Implementations); i++ {
		impl := params.Implementations[i]
		switch impl.Type {
		case types.GoImplementationType:
			goExtractor := GoExtractor{}
			testNames, err := goExtractor.TestNames(impl.Repo.Dir)
			if err != nil {
				return nil, err
			}
			result[types.GoImplementationType] = types.Implementation{
				Repo:      impl.Repo,
				Type:      types.GoImplementationType,
				TestNames: testNames,
			}

		case types.RefImplementationType:
			refExtractor := RefExtractor{}
			testNames, err := refExtractor.TestNames(impl)
			if err != nil {
				return nil, err
			}
			result[types.RefImplementationType] = types.Implementation{
				Repo:      impl.Repo,
				Type:      types.GoImplementationType,
				TestNames: testNames,
			}

		default:
			return result, fmt.Errorf("unknown implementation type: %v", impl.Type)
		}
	}

	return result, nil
}
