package puller

import (
	"os"

	"github.com/go-git/go-git/v5"

	"graphql-go/compatibility-unit-tests/types"
)

type Puller struct {
}

type PullerResult struct {
}

type PullerParams struct {
	Implementation    types.Implementation
	RefImplementation types.Implementation
}

func (p *Puller) Pull(params *PullerParams) (*PullerResult, error) {
	repos := []types.Implementation{
		params.Implementation,
		params.RefImplementation,
	}

	for _, r := range repos {
		if _, err := git.PlainClone("./repos", false, &git.CloneOptions{
			URL:      r.Repo.URL,
			Progress: os.Stdout,
		}); err != nil {
			return nil, err
		}
	}

	return &PullerResult{}, nil
}
