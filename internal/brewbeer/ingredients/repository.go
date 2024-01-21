package ingredients

//go:generate mockery --name=Repository --structname=RepositoryMock --case underscore --output repositorymocks  --outpkg repositorymocks
type Repository interface {
	GetMalt(idMalt string) (*Malt, error)
	GetHop(idHop string) (*Hop, error)
	GetYeast(idYeast string) (*Yeast, error)
}
