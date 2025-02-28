package extractor

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type Extractor struct {
}

type ExtractorResult struct {
}

type ExtractorParams struct {
	RootDir string
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	files := []string{}

	walk := func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			fmt.Println(s)
		}

		files = append(files, s)

		return nil
	}

	filepath.WalkDir(params.RootDir, walk)

	return nil, nil
}
