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
func (abv *abvImpl) Estimate(params Params) (float64, error) {
	result := abv.service.CalculateABV(&params.Fermentation)

	return result, nil
}
