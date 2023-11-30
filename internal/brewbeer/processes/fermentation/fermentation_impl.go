package fermentation

import (
	"eccea/internal/brewbeer/ingredients"
)

type service struct {
	repo ingredients.Repository
}

func NewService(repo ingredients.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Do(params *Params) *Results {
	return nil
}
