package extractor

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"

	"github.com/evanw/esbuild/pkg/api"
)

type Extractor struct {
}

type ExtractorResult struct {
	TestFiles []string
}

type ExtractorParams struct {
	Source string
}

func (e *Extractor) readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (e *Extractor) fileContentToFileAST(fileContent string) (*js.AST, error) {
	ast, err := js.Parse(parse.NewInputString(fileContent), js.Options{})
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

	return ast, nil
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
		fileContent, err := e.readFile(testFile)
		if err != nil {
			return nil, err
		}

		if testFile == "repos/graphql-graphql-js/src/execution/__tests__/schema-test.ts" {
			transformResult := api.Transform(string(fileContent), api.TransformOptions{
				Loader:            api.LoaderTSX,
				MinifyWhitespace:  false,
				MinifyIdentifiers: false,
				MinifySyntax:      false,
			})

			if len(transformResult.Errors) > 0 {
				errs := errors.New("")
				for _, err := range transformResult.Errors {
					fmt.Errorf("%v: %w", err, errs)
					return nil, errs
				}
			}

			ast, err := e.fileContentToFileAST(string(transformResult.Code))
			if err != nil {
				return nil, err
			}

			js.Walk(&walker{}, ast)
		}
	}

	return &ExtractorResult{}, nil
}

type walker struct{}

func (w *walker) Enter(n js.INode) js.IVisitor {
	switch n := n.(type) {
	case *js.Var:
		log.Println(string(n.Data))
	}

	return w
}

func (w *walker) Exit(n js.INode) {
}
