package dto

import "eccea/internal/brewbeer/ingredients"

type HopResponse struct {
	ID              string                   `json:"id"`
	Name            string                   `json:"name"`
	AlphaAcids      float64                  `json:"alpha_acids"`
	BetaAcids       float64                  `json:"beta_acids"`
	Characteristics []CharacteristicResponse `json:"characteristics"`
}

func (h *HopResponse) ToDomain() *ingredients.Hop {
	ccharacts := []ingredients.Characteristic{}
	for _, cc := range h.Characteristics {
		ccharact := ingredients.Characteristic{
			TypeAdjetives:      cc.AdjetiveID,
			CharacteristicType: cc.Type,
			Description:        cc.Description,
			Contribution:       cc.Contribution,
		}
		ccharacts = append(ccharacts, ccharact)
	}

	return ingredients.NewHopBuilder().
		Name(h.Name).
		Characteristics(ccharacts).
		AlphaAcids(h.AlphaAcids).
		BetaAcids(h.BetaAcids).
		Build()
}
