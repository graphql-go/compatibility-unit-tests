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
		name := "./repos/" + r.Repo.Name
		if _, err := os.Stat(name); os.IsNotExist(err) {
			if err := os.Mkdir(name, os.ModePerm); err != nil {
				return nil, err
			}
		}
		if _, err := git.PlainClone(name, false, &git.CloneOptions{
			URL:      r.Repo.URL,
			Progress: os.Stdout,
		}); err != nil {
			return nil, err
		}
	}

	return &PullerResult{}, nil
}
