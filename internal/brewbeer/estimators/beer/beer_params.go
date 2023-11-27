package beer

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/fermentation"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

type Params struct {
	Maceration   maceration.Params
	Boiling      boiling.Params
	Fermentation fermentation.Params
	Maturation   maturation.Params
}

func NewParams(maceration maceration.Params,
	boiling boiling.Params,
	fermentation fermentation.Params,
	maturation maturation.Params) *Params {
	return &Params{
		Maceration:   maceration,
		Boiling:      boiling,
		Fermentation: fermentation,
		Maturation:   maturation,
	}
}
