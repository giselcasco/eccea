package color

import (
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

type Params struct {
	Maceration maceration.Params
	Maturation maturation.Params
}

func NewParams(maceration maceration.Params, maturation maturation.Params) *Params {
	return &Params{
		Maturation: maturation,
		Maceration: maceration,
	}
}
