package responses

import (
	"eccea/internal/brewbeer/ingredients"
)

type Hop struct {
	ID         string
	AlphaAcids float64
	BetaAcids  float64
}

func (h Hop) ToDomain() *ingredients.Hop {
	hop := &ingredients.Hop{}
	hop.SetID(h.ID)
	hop.SetAlphaAcids(h.AlphaAcids)
	hop.SetBetaAcids(h.BetaAcids)
	return hop
}
