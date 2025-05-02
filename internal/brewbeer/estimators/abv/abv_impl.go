package abv

import "eccea/internal/brewbeer/processes/fermentation"

// abvImpl es la implementacion para el caso de uso de calculo de ABV
type abvImpl struct {
	service fermentation.Service
}

func NewABVImpl(service fermentation.Service) ABV {
	return &abvImpl{
		service: service,
	}
}

// implementación para el calculo de la estimación del ABV
func (abv *abvImpl) EstimateAlcoholByVolume(params Params) (float64, error) {
	return abv.service.CalculateAlcoholByVolume(params.Fermentation.InitialDensity, params.Fermentation.FinalDensity)
}
