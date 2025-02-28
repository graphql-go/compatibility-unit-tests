package extractor

import (
	"io/fs"
	"log"
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

	log.Println(testFiles)

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
