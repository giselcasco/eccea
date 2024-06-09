package ibu

import (
	"eccea/internal/brewbeer/processes/boiling"
)

/* 
    Params contiene los parametros necesarios para el calculo del IBU,
    en este caso contiene los parametros asociados a la cocción
*/ 
type Params struct {
	Boiling boiling.Params
}

func NewParams(boiling boiling.Params) *Params {
	return &Params{
		Boiling: boiling,
	}
}
