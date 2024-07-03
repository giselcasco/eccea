package maceration

import (
	"eccea/internal/brewbeer/ingredients"
)

type service struct {
	repo ingredients.Repository
}

func NewService(repo ingredients.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) EstimateColor(params *Params) *ColorResults {
	results := NewColorResults()
	finalColorSRM := s.calculateFinalColorSRM(params)
	if finalColorSRM > 0 {
		results.SetColor(uint64(finalColorSRM))
	}
	return results
}

func (*service) calculateFinalColorSRM(params *Params) float64 {
	var sum float64
	for _, m := range params.MaltAdditions {
		sum = +(m.Quantity * m.ColorSRM)
	}

	if params.WortAmount > 0 {
		color := sum / (params.WortAmount * 0.96)
		return color
	}
	return 0
}
