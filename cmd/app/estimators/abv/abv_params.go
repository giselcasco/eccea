package abv

import (
	"eccea/internal/brewbeer/processes/fermentation"
)

type Params struct {
	Fermentation fermentation.Params
}

func NewParams(fermentation fermentation.Params) *Params {
	return &Params{
		Fermentation: fermentation,
	}
}
