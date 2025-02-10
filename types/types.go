package types

type Repository struct {
	Name string
	URL  string
}

type Implementation struct {
	Repo Repository
}
