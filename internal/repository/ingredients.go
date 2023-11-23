package repository

import "eccea/internal/brewbeer/ingredients"

type Ingredients interface {
	GetMalt(idMalt string) (*ingredients.Malt, error)
	GetHop(idHop string) (*ingredients.Hop, error)
	GetYeast(idYeast string) (*ingredients.Yeast, error)
}

type ingredientsRepository struct {
}

func NewIngredientsRepository() Ingredients {
	return &ingredientsRepository{}
}

func (i *ingredientsRepository) GetMalt(idMalt string) (*ingredients.Malt, error) {
	return nil, nil
}
func (i *ingredientsRepository) GetHop(idHop string) (*ingredients.Hop, error) {
	return nil, nil
}
func (i *ingredientsRepository) GetYeast(idYeast string) (*ingredients.Yeast, error) {
	return nil, nil
}
