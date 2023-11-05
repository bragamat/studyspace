package repo

type Repository struct{}

type RepositoryStats interface {
	GetStats(url string) *Repository
}
