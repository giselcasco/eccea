package fermentation

type service struct {
}

func NewService() Service {
	return &service{}
}

// CalculateABV es el método que calcula el abv a partir de los valores en los parametros
func (s *service) CalculateABV(params *Params) float64 {
	return (params.InitialDensity - params.FinalDensity) * 131.25
}

// EstimateBeer es el método que calcula el abv a partir de los valores en los parametros
// y devuelve las caracteristicas que aporta la levadura a la cerveza
func (s *service) EstimateBeer(params *Params) (*Results, error) {
	abv := (params.InitialDensity - params.FinalDensity) * 131.25

	// TODO
	results := &Results{}
	results.SetABV(abv)
	return results, nil
}
