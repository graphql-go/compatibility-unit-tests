package extractor

import (
  "github.com/robertkrimen/otto/parser"
)

type Extractor struct {
}

type ExtractorResult struct {
}

type ExtractorParams struct {
  Source string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	if _, err := parser.ParseFile(nil, "", params.Source, 0); err != nil {
    return nil, err
  }

	return &ExtractorResult{}, nil
}
