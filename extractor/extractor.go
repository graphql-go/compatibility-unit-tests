package extractor

import (
	"fmt"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"
)

type Extractor struct {
}

type ExtractorResult struct {
}

type ExtractorParams struct {
	Source string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	ast, err := js.Parse(parse.NewInputString(params.Source), js.Options{})
	if err != nil {
		return nil, err
	}

	for _, a := range ast.List {
		str, ok := a.(*js.VarDecl)
		if ok {
			fmt.Printf("default: %+v\n", str.List[0].Default)
			fmt.Printf("binding: %+v\n", str.List[0].Binding)
		}
	}

	return &ExtractorResult{}, nil
}
