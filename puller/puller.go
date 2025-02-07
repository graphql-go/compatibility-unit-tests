package puller

import (
	"os"

	"github.com/go-git/go-git/v5"
)

type Puller struct {
}

type PullerResult struct {
}

type PullerParams struct {
	RepoURL string
}

func (p *Puller) Pull(params *PullerParams) (*PullerResult, error) {
	if _, err := git.PlainClone("./repos", false, &git.CloneOptions{
		URL:      params.RepoURL,
		Progress: os.Stdout,
	}); err != nil {
		return nil, err
	}

	return &PullerResult{}, nil
}
