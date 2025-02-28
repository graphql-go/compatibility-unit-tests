package extractor

import (
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Extractor struct {
}

type ExtractorResult struct {
	TestFiles []string
}

type ExtractorParams struct {
	RootDir string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {

	testFiles, err := e.testFiles(params.RootDir)
	if err != nil {
		return nil, err
	}

	funcNames, err := e.funcNames(testFiles)
	if err != nil {
		return nil, err
	}

	log.Println(funcNames)

	return &ExtractorResult{
		TestFiles: testFiles,
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

func (e *Extractor) funcNames(testFiles []string) ([]string, error) {
	result := []string{}

	filePath := testFiles[0]

	goFile, err := os.Open(filePath)
	if err != nil {
		return result, err
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", goFile, parser.ParseComments)
	if err != nil {
		return result, nil
	}

	log.Println(filePath)
	log.Println(astFile.Decls)

	return []string{}, nil
}
