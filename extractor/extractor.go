package extractor

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

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

func (e *Extractor) readFiles() (*ExtractorResult, error) {
	rootDir := "./repos/graphql-graphql-js"

	testsFiles := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		files, err := ioutil.ReadDir(s)
		if err != nil {
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				if strings.HasSuffix(file.Name(), "-test.ts") {
					testsFiles = append(testsFiles, file.Name())
				}
			}
		}

		return nil
	}

	filepath.WalkDir(rootDir, walk)

	return nil, nil
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	if _, err := e.readFiles(); err != nil {
		return nil, err
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
