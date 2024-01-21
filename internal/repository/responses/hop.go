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
	AlphaAcids float32
	BetaAcids  float32
}

func (h Hop) ToDomain(idHop string) *ingredients.Hop {
	hop := &ingredients.Hop{}
	hop.SetID(idHop)
	hop.SetAlphaAcids(h.AlphaAcids)
	hop.SetBetaAcids(h.BetaAcids)
	return hop
}
