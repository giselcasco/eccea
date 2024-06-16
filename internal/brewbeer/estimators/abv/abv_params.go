package abv

import (
	"eccea/internal/brewbeer/processes/fermentation"
)

/* 
    Params contiene los parametros necesarios para el calculo del ABV,
    en este caso contiene los parametros asociados a la fermentación
*/ 
type Params struct {
	Fermentation fermentation.Params
}

func NewParams(fermentation fermentation.Params) *Params {
	return &Params{
		Fermentation: fermentation,
	}
}
