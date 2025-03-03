package types

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

type Implementation struct {
	Repo              Repository
	Type              ImplementationType
	TestNames         []string
	TestNamesFilePath string
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
