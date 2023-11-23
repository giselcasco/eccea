package maceration

import "eccea/internal/repository"

type service struct {
	ingredients repository.Ingredients
}

func NewService(ingredients repository.Ingredients) Maceration {
	return &service{
		ingredients: ingredients,
	}
}

func (s *service) Do(params *Params) *Results {
	return nil
}
