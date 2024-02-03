package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"math"
	"strings"
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

	for _, hopAddition := range params.HopAdditions {
		hop, errRepo := s.repo.GetHop(hopAddition.ID)
		if errRepo != nil {
			return nil, errRepo
		}
		hops = append(hops, *hop)
	}

	if estimateFunc, ok := funcByUseCaseKey[useCaseKey]; ok {
		estimateFunc(params, hops, estimateResults)
	}

	return estimateResults, nil
}

type estimate func(params *Params, hops []ingredients.Hop, result *Results)

var funcByUseCaseKey = map[string]estimate{
	"ibu": estimateIBU,
}

func estimateIBU(params *Params, hops []ingredients.Hop, result *Results) {
	for _, hopAddition := range params.HopAdditions {
		if hop := getHop(hops, hopAddition.ID); hop != nil {
			result.ibu += calculateIBU(params, hopAddition, *hop)
		}
	}
}

func getHop(hops []ingredients.Hop, idHop string) *ingredients.Hop {
	for _, hop := range hops {
		if strings.EqualFold(hop.ID(), idHop) {
			return &hop
		}
	}
	return nil
}

func calculateIBU(params *Params, addition HopAdditions, hop ingredients.Hop) float64 {
	firstFactor := greatnessFactor(params.InitialDensity) * boilingTimeFactor(addition.TimeOfWork)
	secondFactor := proportionOfAlphaAcidUsed(hop.AlphaAcids(), addition.Quantity, params.WortAmount)

	return firstFactor * float64(secondFactor)
}

func greatnessFactor(initialDensity uint32) float64 {
	base := 0.000125
	exponent := (initialDensity / 100) - 1
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

	if divisor > 0 {
		return dividend / divisor
	}
	return 0
}
