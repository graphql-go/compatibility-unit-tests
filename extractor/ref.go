package extractor

import (
	"os"
	"strings"

	"graphql-go/compatibility-unit-tests/types"
)

type RefExtractor struct{}

func (e *RefExtractor) TestNames(impl types.Implementation) ([]string, error) {
	f, err := os.ReadFile(impl.TestNamesFilePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(f), "\n"), nil
}
