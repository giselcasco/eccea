package repository

import "eccea/internal/brewbeer/ingredients"

//go:generate mockery --name=Repository --structname=RepositoryMock --output repositorymocks  --outpkg repositorymocks
type Reader interface {
	GetMaltByName(maltName string) (*ingredients.Malt, error)
	GetHopByName(hopName string) (*ingredients.Hop, error)
	GetYeastByName(yeastName string) (*ingredients.Yeast, error)
}
