package extractor

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path"
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

	testFiles := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		testsFiles = append(testsFiles, s)
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
					path := path.Join(s, file.Name())
					testFiles = append(testFiles, path)
				}
			}
		}

		return nil
	}

	filepath.WalkDir(rootDir, walk)

	return &ExtractorResult{
		TestFiles: testFiles,
	}, nil
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	extractorResult, err := e.readFiles()
	if err != nil {
		return nil, err
	}

	for _, testFile := range extractorResult.TestFiles {
		fmt.Printf("testFile: %+v\n", testFile)
		if testFile == "repos/graphql-graphql-js/src/execution/__tests__/schema-test.ts" {
			log.Println(testFile)
		}
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
