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

const useCaseKey = "abv"

// implementación para el calculo de la estimación del ABV
func (abv *abvImpl) Estimate(params Params) (Estimation, error) {
	result := abv.service.Do(&params.Fermentation, useCaseKey)

	return Estimation{
		ABV: result.ABV(),
	}, nil
}
