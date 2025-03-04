package types

import "fmt"

const taggedRepoURL string = "%s/releases/tag/%s"

type ImplementationType uint

const (
	GoImplementationType = iota + 1
	RefImplementationType
)

type Repository struct {
	Name          string
	URL           string
	ReferenceName string
	Dir           string
}

func (r *Repository) String() string {
	base := fmt.Sprintf("Implementation: %s\n", taggedRepoURL)
	return fmt.Sprintf(base, r.URL, r.ReferenceName)
}

type Implementation struct {
	Repo              Repository
	Type              ImplementationType
	TestNames         []string
	TestNamesFilePath string
}

func (i *Implementation) MapKey() string {
	return i.Repo.String()
}

type ImplementationTest struct {
	Name string
}

type SuccessfulTest struct {
	Name string
}

type FailedTest struct {
	Name string
}
