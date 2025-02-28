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
	Repo Repository
}

type ImplementationTest struct {
	Name string
}
