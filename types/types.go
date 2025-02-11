package types

type Repository struct {
	Name          string
	URL           string
	ReferenceName string
}

type Implementation struct {
	Repo Repository
}
