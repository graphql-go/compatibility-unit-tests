package extractor

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"
)

type Extractor struct {
}

type ExtractorResult struct {
	TestFiles []string
}

type ExtractorParams struct {
	Source string
}

func (e *Extractor) readFiles() (*ExtractorResult, error) {
	rootDir := "./repos/graphql-graphql-js"

	testsFiles := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		testsFiles = append(testsFiles, s)

		return nil
	}

	filepath.WalkDir(rootDir, walk)

	return &ExtractorResult{
		TestFiles: testsFiles,
	}, nil
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	extractorResult, err := e.readFiles()
	if err != nil {
		return nil, err
	}

	for _, testFile := range extractorResult.TestFiles {
		fmt.Printf("testFile: %+v\n", testFile)
	}

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
