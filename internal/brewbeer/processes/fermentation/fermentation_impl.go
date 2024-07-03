package fermentation

type service struct {
}

func NewService() Service {
	return &service{}
}

// estimateABV es el método que calcula el abv a partir de los valores en los parametros
func (s *service) CalculateABV(params *Params) float64 {
	return (params.InitialDensity - params.FinalDensity) * 131.25
}
