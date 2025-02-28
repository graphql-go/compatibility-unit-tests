package extractor

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"graphql-go/compatibility-unit-tests/types"
)

type Extractor struct {
}

type ExtractorParams struct {
	Implementation    types.Implementation
	RefImplementation types.Implementation
}

type ExtractorResult struct {
	TestNames map[types.ImplementationType]types.Implementation
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	testNames, err := e.implementationTestNames(ImplTestNamesParams{
		Implementations: []types.Implementation{params.Implementation, params.RefImplementation},
	})
	if err != nil {
		return nil, err
	}

	return &ExtractorResult{
		TestNames: testNames,
	}, nil
}

type ImplTestNamesParams struct {
	Implementations []types.Implementation
}

func (e *Extractor) implementationTestNames(params ImplTestNamesParams) (map[types.ImplementationType]types.Implementation, error) {
	result := map[types.ImplementationType]types.Implementation{}

	for i := 0; i < len(params.Implementations); i++ {
		impl := params.Implementations[i]
		switch impl.Type {
		case types.GoImplementationType:
			testNames, err := e.goTestNames(impl.Repo.Dir)
			if err != nil {
				return nil, err
			}
			result[types.GoImplementationType] = types.Implementation{
				Repo:      impl.Repo,
				Type:      types.GoImplementationType,
				TestNames: testNames,
			}

		case types.RefImplementationType:
			testNames, err := e.refTestNames(impl.Repo.Dir)
			if err != nil {
				return nil, err
			}
			result[types.RefImplementationType] = types.Implementation{
				Repo:      impl.Repo,
				Type:      types.GoImplementationType,
				TestNames: testNames,
			}

		default:
			return result, fmt.Errorf("unknown implementation type: %v", impl.Type)
		}
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

func (e *Extractor) refTestNames(refRootDir string) ([]string, error) {
	f, err := os.ReadFile(refRootDir)
	if err != nil {
		return nil, err
	}

	log.Println(f)

	return []string{}, nil
}
