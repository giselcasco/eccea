package ibu

import "eccea/internal/brewbeer/processes/boiling"

type ibuImpl struct {
	service boiling.Service
}

func NewIBUImpl(service boiling.Service) IBU {
	return &ibuImpl{
		service: service,
	}
}

const useCaseKey = "ibu"

func (ibu *ibuImpl) Estimate(params Params) (Estimation, error) {
	result, err := ibu.service.Do(&params.Boiling, useCaseKey)
	if err != nil {
		return Estimation{}, err
	}

	return Estimation{
		IBU: result.IBU(),
	}, nil
}
