package ibu

import (
	"eccea/internal/brewbeer/processes/boiling"
)

type Params struct {
	Boiling boiling.Params
}

func NewParams(boiling boiling.Params) *Params {
	return &Params{
		Boiling: boiling,
	}
}
