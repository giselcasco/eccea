package flavor

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

type Params struct {
	Maceration maceration.Params
	Boiling    boiling.Params
	Maturation maturation.Params
}

func NewParams(maceration maceration.Params, boiling boiling.Params, maturation maturation.Params) *Params {
	return &Params{
		Maceration: maceration,
		Boiling:    boiling,
		Maturation: maturation,
	}
}
