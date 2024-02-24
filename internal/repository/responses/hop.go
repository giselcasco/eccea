package responses

import (
	"eccea/internal/brewbeer/ingredients"
)

type Hop struct {
	/*
		FlavorCharacteristics    []characteristic.Flavor
		SmellCharacteristics     []characteristic.Smell
		MouthfeelCharacteristics []characteristic.Mouthfeel
	*/
	AlphaAcids float64
	BetaAcids  float64
}

func (h Hop) ToDomain(idHop string) *ingredients.Hop {
	hop := &ingredients.Hop{}
	hop.SetID(idHop)
	hop.SetAlphaAcids(h.AlphaAcids)
	hop.SetBetaAcids(h.BetaAcids)
	return hop
}
