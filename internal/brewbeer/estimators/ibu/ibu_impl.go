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

// implementación para el calculo de la estimación del IBU
func (ibu *ibuImpl) Estimate(params Params) (float64, error) {
	result, err := ibu.service.EstimateIBU(&params.Boiling)
	if err != nil {
		return 0, err
	}

	return result, nil
}
