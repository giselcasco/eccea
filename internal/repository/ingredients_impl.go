package repository

import "eccea/internal/brewbeer/ingredients"

type ingredientRepo struct {
}

func NewIngredientsRepository() ingredients.Repository {
	return &ingredientRepo{}
}

func (i *ingredientRepo) GetMalt(idMalt string) (*ingredients.Malt, error) {
	return nil, nil
}

func (i *ingredientRepo) GetHop(idHop string) (*ingredients.Hop, error) {
	return nil, nil
}

func (i *ingredientRepo) GetYeast(idYeast string) (*ingredients.Yeast, error) {
	return nil, nil
}
