package ibu

import "eccea/internal/brewbeer/processes/boiling"

// ibuImpl es la implementacion para el caso de uso de calculo de IBU
type ibuImpl struct {
	service boiling.Service
}

func NewIBUImpl(service boiling.Service) IBU {
	return &ibuImpl{
		service: service,
	}
}

const useCaseKey = "ibu"

// implementación para el calculo de la estimación del IBU
func (ibu *ibuImpl) Estimate(params Params) (Estimation, error) {
	result, err := ibu.service.Do(&params.Boiling, useCaseKey)
	if err != nil {
		return Estimation{}, err
	}

	return Estimation{
		IBU: result.IBU(),
	}, nil
}
