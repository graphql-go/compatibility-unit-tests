package extractor

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"graphql-go/compatibility-unit-tests/types"
)

type Extractor struct {
}

type ExtractorParams struct {
	ImplementationType    types.ImplementationType
	RefImplementationType types.ImplementationType
	RootDir               string
}

type ExtractorResult struct {
	ReferenceTestNames      []string
	ImplementationTestNames []string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	refTestNames, err := e.implementationTestNames(params.RefImplementationType, params.RootDir)
	if err != nil {
		return nil, err
	}

	implTestNames, err := e.implementationTestNames(params.ImplementationType, params.RootDir)
	if err != nil {
		return nil, err
	}

	return &ExtractorResult{
		ReferenceTestNames:      refTestNames,
		ImplementationTestNames: implTestNames,
	}, nil
}

func (e *Extractor) implementationTestNames(implementationType types.ImplementationType, rootDir string) ([]string, error) {
	result := []string{}

	switch implementationType {
	case types.GoImplementationType:
		testNames, err := e.goTestNames(rootDir)
		if err != nil {
			return nil, err
		}
		result = append(result, testNames...)
	case types.RefImplementationType:
		testNames, err := e.refTestNames(rootDir)
		if err != nil {
			return nil, err
		}
		result = append(result, testNames...)
	default:
		return []string{}, fmt.Errorf("unknown implementation type: %v", implementationType)
	}

	return result, nil
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

func (e *Extractor) goTestNames(rootDir string) ([]string, error) {
	testFiles, err := e.testFiles(rootDir)
	if err != nil {
		return nil, err
	}

	testNames, err := e.testNames(testFiles)
	if err != nil {
		return nil, err
	}

	return testNames, nil
}

func (e *Extractor) refTestNames(rootDir string) ([]string, error) {
	return []string{}, nil
}
