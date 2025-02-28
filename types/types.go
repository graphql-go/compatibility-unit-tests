package types

type Repository struct {
	Name          string
	URL           string
	ReferenceName string
	Dir           string
}

type Implementation struct {
	Repo Repository
}
