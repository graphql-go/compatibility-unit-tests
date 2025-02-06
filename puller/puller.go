package puller

import (
	"log"
	"os"

	"github.com/google/go-github/v69/github"
)

type Puller struct {
}

type PullerResult struct {
}

type PullerParams struct {
}

func (p *Puller) Pull(params *PullerParams) (*PullerResult, error) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return nil, nil
	}

	client := github.NewClient(nil).WithAuthToken(token)

	log.Printf("client: %+v", client)

	return &PullerResult{}, nil
}
