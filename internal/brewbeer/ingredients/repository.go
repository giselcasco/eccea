package ingredients

//go:generate mockery --name=Repository --structname=RepositoryMock --output repositorymocks  --outpkg repositorymocks
type Reader interface {
	GetMaltByName(maltName string) (*Malt, error)
	GetHopByName(hopName string) (*Hop, error)
	GetYeastByName(yeastName string) (*Yeast, error)
}
