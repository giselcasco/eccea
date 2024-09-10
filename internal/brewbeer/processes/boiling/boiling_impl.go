package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"math"
	"strings"
)

// service es el implementador de los metodos de proceso de coccion
type service struct {
	repo ingredients.Reader
}

func NewService(repo ingredients.Reader) Service {
	return &service{
		repo: repo,
	}
}

/*
EstimateIBU es el metodo que calcula el ibu a partir de los valores en los parametros
ingresados por el usuario.
*/
func (s *service) EstimateIBU(params *Params) (float64, error) {
	var ibu float64
	for _, hopAddition := range params.HopAdditions {
		ibu += calculateIBU(params, hopAddition)
	}

	return ibu, nil
}

// getHop busca el lupudo en la lista de lupulos "hops" cuyo ID corresponda con "idHop"
// opcional para busqueda en db del alphaAcids del hop
func getHop(hops []ingredients.Hop, idHop string) *ingredients.Hop {
	for _, hop := range hops {
		if strings.EqualFold(hop.Name(), idHop) {
			return &hop
		}
	}
	return nil
}

// calculateIBU implementa la formula de calculo del IBU de Glenn Tinseth
func calculateIBU(params *Params, addition Hop) float64 {
	firstFactor := greatnessFactor(params.InitialDensity)
	secondFactor := boilingTimeFactor(addition.TimeOfWork)
	thirdFactor := proportionOfAlphaAcidUsed(addition.AlphaAcids, addition.Quantity)
	divisor := float64(params.WortAmount) * 4.15

	if divisor > 0 {
		return (firstFactor * secondFactor * thirdFactor) / divisor
	}
	return 0
}

func greatnessFactor(initialDensity float64) float64 {
	if initialDensity <= 0 {
		return 0
	}

	firstFactor := 1.65
	base := 0.000125
	exponent := initialDensity - 1.0
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
