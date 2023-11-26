package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/repository"
	"math"
)

type service struct {
	ingredients repository.Ingredients
}

func NewService(ingredients repository.Ingredients) Boiling {
	return &service{
		ingredients: ingredients,
	}
}

func (s *service) Do(params *Params) (*Results, error) {
	estimateResults := &Results{}
	var ibus float64
	for _, hopAddition := range params.Additions() {
		hop, err := s.ingredients.GetHop(hopAddition.id)
		if err != nil {
			return nil, err
		}
		ibus += s.calculateIBU(params, hopAddition, hop)
	}

	estimateResults.ibu = ibus

	return nil, nil
}

func (s *service) calculateIBU(params *Params, addition HopAdditions, hop *ingredients.Hop) float64 {
	firstFactor := s.greatnessFactor(params.initialDensity) * s.boilingTimeFactor(addition.timeOfWork)
	secondFactor := s.proportionOfAlphaAcidUsed(hop.AlphaAcids(), addition.quantity, params.volume)

	return firstFactor * float64(secondFactor)
}

func (s *service) greatnessFactor(initialDensity uint32) float64 {
	base := 0.000125
	exponent := initialDensity - 1
	firstFactor := 1.65
	secondFactor := math.Pow(base, float64(exponent))

	return firstFactor * secondFactor
}

func (s *service) boilingTimeFactor(timeOfWork uint32) float64 {
	exponent := -0.04 * float64(timeOfWork)
	eRaisedExp := math.Exp(exponent)
	dividend := 1 - eRaisedExp
	divisor := 4.15

	return dividend / divisor
}

func (s *service) proportionOfAlphaAcidUsed(alphaAcids float32, quantity float32, volume uint32) float32 {
	dividend := alphaAcids * quantity * 1000
	divisor := float32(volume)

	return dividend / divisor
}
