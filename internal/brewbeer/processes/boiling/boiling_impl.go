package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"math"
)

type service struct {
	repo ingredients.Repository
}

func NewService(repo ingredients.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Do(params *Params, useCaseKey string) (*Results, error) {
	var hops []ingredients.Hop
	estimateResults := &Results{}

	for _, hopAddition := range params.Additions() {
		hop, errRepo := s.repo.GetHop(hopAddition.id)
		if errRepo != nil {
			return nil, errRepo
		}
		hops = append(hops, *hop)
	}

	if estimateFunc, ok := funcByUseCaseKey[useCaseKey]; ok {
		errEstimate := estimateFunc(params, hops, estimateResults)
		if errEstimate != nil {
			return nil, errEstimate
		}
	}

	return estimateResults, nil
}

type estimate func(params *Params, hops []ingredients.Hop, result *Results) error

var funcByUseCaseKey = map[string]estimate{
	"ibu": estimateIBU,
}

func estimateIBU(params *Params, hops []ingredients.Hop, result *Results) error {
	var ibus float64
	for _, hopAddition := range params.Additions() {
		if hop := getHop(hops, hopAddition.id); hop != nil {
			ibus += calculateIBU(params, hopAddition, *hop)
		}
	}

	result.ibu = ibus

	return nil
}

func getHop(hops []ingredients.Hop, idHop string) *ingredients.Hop {
	for _, hop := range hops {
		if hop.ID() == idHop {
			return &hop
		}
	}
	return nil
}

func calculateIBU(params *Params, addition HopAdditions, hop ingredients.Hop) float64 {
	firstFactor := greatnessFactor(params.initialDensity) * boilingTimeFactor(addition.timeOfWork)
	secondFactor := proportionOfAlphaAcidUsed(hop.AlphaAcids(), addition.quantity, params.volume)

	return firstFactor * float64(secondFactor)
}

func greatnessFactor(initialDensity uint32) float64 {
	base := 0.000125
	exponent := initialDensity - 1
	firstFactor := 1.65
	secondFactor := math.Pow(base, float64(exponent))

	return firstFactor * secondFactor
}

func boilingTimeFactor(timeOfWork uint32) float64 {
	exponent := -0.04 * float64(timeOfWork)
	eRaisedExp := math.Exp(exponent)
	dividend := 1 - eRaisedExp
	divisor := 4.15

	return dividend / divisor
}

func proportionOfAlphaAcidUsed(alphaAcids float32, quantity float32, volume uint32) float32 {
	dividend := alphaAcids * quantity * 1000
	divisor := float32(volume)

	return dividend / divisor
}
