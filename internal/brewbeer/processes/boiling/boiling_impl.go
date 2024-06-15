package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"fmt"
	"math"
	"strings"
)

// service es el implementador de los metodos de proceso de coccion 
type service struct {
	repo ingredients.Repository
}

func NewService(repo ingredients.Repository) Service {
	return &service{
		repo: repo,
	}
}

/* 
    Do es la implementaciòn del proceso de coccion
    el mismo consiste 
*/
func (s *service) Do(params *Params, useCaseKey string) (*Results, error) {
	var hops []ingredients.Hop
	var estimateResults Results

	for _, hopAddition := range params.HopAdditions {
		hop, errRepo := s.repo.GetHop(hopAddition.ID)
		if errRepo != nil {
			return &estimateResults, errRepo
		}
		hops = append(hops, *hop)
	}

	if estimateFunc, ok := funcByUseCaseKey[useCaseKey]; ok {
		estimateFunc(params, hops, &estimateResults)
	}

	return &estimateResults, nil
}

type estimate func(params *Params, hops []ingredients.Hop, result *Results)

var funcByUseCaseKey = map[string]estimate{
	"ibu": estimateIBU,
}

func estimateIBU(params *Params, hops []ingredients.Hop, result *Results) {
	var ibu float64
	for _, hopAddition := range params.HopAdditions {
		if hop := getHop(hops, hopAddition.ID); hop != nil {
			ibu += calculateIBU(params, hopAddition, *hop)
		}
	}

	if ibu > 0 {
		result.ibu = fmt.Sprintf("%.1f", ibu)
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
	firstFactor := greatnessFactor(params.InitialDensity)
	secondFactor := boilingTimeFactor(addition.TimeOfWork)
	thirdFactor := proportionOfAlphaAcidUsed(hop.AlphaAcids(), addition.Quantity)
	divisor := float64(params.WortAmount) * 4.15

	return (firstFactor * secondFactor * thirdFactor) / divisor
}

func greatnessFactor(initialDensity uint64) float64 {
	firstFactor := 1.65

	base := 0.000125
	exponent := (float64(initialDensity) / 1000.0) - 1.0
	secondFactor := math.Pow(base, exponent)

	return firstFactor * secondFactor
}

func boilingTimeFactor(timeOfWork uint64) float64 {
	exponent := -0.04 * float64(timeOfWork)
	eRaisedExp := math.Exp(exponent)
	dividend := 1 - eRaisedExp

	return dividend
}

func proportionOfAlphaAcidUsed(alphaAcids float64, quantity float64) float64 {
	return (alphaAcids / 100.0) * quantity * 1000
}
