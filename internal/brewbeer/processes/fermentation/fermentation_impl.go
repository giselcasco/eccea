package fermentation

import "errors"

type service struct {
}

func NewService() Service {
	return &service{}
}

// CalculateAlcoholByVolume calcula el porcentaje de alcohol basado en la densidad inicial y final.
func (s *service) CalculateAlcoholByVolume(initialDensity, finalDensity float64) (float64, error) {
	if initialDensity <= 0 || finalDensity <= 0 {
		return 0, errors.New("Both densities must be greater than zero")
	}
	if finalDensity > initialDensity {
		return 0, errors.New("final density cannot be greater than initial density")
	}

	abv := (initialDensity - finalDensity) * 131.25
	return abv, nil
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
