package extractor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Extractor struct {
}

type ExtractorParams struct {
	RootDir string
}

type ExtractorResult struct {
	TestNames []string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {

	testFiles, err := e.testFiles(params.RootDir)
	if err != nil {
		return nil, err
	}

	testNames, err := e.testNames(testFiles)
	if err != nil {
		return nil, err
	}

	return &ExtractorResult{
		TestNames: testNames,
	}, nil
}

func (e *Extractor) testFiles(rootDir string) ([]string, error) {
	testFiles := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(s, "_test.go") {
			testFiles = append(testFiles, s)
		}

		return nil
	}

	filepath.WalkDir(rootDir, walk)

	return testFiles, nil
}

func (e *Extractor) readFile(filePath string) (*os.File, error) {
	goFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return goFile, nil
}

func (e *Extractor) readFuncNames(filePath string) ([]string, error) {
	goFile, err := e.readFile(filePath)
	if err != nil {
		return nil, err
	}
	defer goFile.Close()

	funcNames := []string{}
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", goFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range astFile.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			funcNames = append(funcNames, t.Name.Name)
		}
	}

	return funcNames, nil
}

func (e *Extractor) testNames(testFiles []string) ([]string, error) {
	result := []string{}

	for _, filePath := range testFiles {
		funcNames, err := e.readFuncNames(filePath)
		if err != nil {
			return result, err
		}

		result = append(result, funcNames...)
	}

	return result, nil
}
