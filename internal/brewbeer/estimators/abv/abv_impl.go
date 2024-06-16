package abv


// abvImpl es la implementacion para el caso de uso de calculo de ABV
type abvImpl struct {
	service fermentation.Service
}

func NewIBUImpl(service fermentation.Service) IBU {
	return &ibuImpl{
		service: service,
	}
}

const useCaseKey = "abv"

// implementación para el calculo de la estimación del ABV
func (abv *abvImpl) Estimate(params Params) (Estimation, error) {
	result, err := ibu.service.Do(&params.Fermentation, useCaseKey)
	if err != nil {
		return Estimation{}, err
	}

	return Estimation{
		ABV: result.ABV(),
	}, nil
}
