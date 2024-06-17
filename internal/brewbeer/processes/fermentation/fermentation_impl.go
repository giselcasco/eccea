package fermentation

import "fmt"

type service struct {
}

func NewService() Service {
	return &service{}
}

/*
	  Do es la implementaciòn del proceso de fermentacion, de este se obtienen características de menor percepción
	y el valor de ABV (alcohol by volume)
	  useCaseKey hace referencia al caso de uso con el que se consulta al proceso,

puede ser "abv", "smell", flavor", "color" or "beer"
*/
func (s *service) Do(params *Params, useCaseKey string) *Results {
	estimateResults := &Results{}
	if estimateFunc, ok := funcByUseCaseKey[useCaseKey]; ok {
		estimateFunc(params, estimateResults)
	}

	return estimateResults
}

type estimate func(params *Params, result *Results)

var funcByUseCaseKey = map[string]estimate{
	"abv": calculateABV,
}

// estimateABV es el metodo que calcula el abv a partir de los valores en los parametros
func calculateABV(params *Params, result *Results) {
	if abv := float32(params.InitialDensity-params.FinalDensity) * 131.25; abv > 0 {
		result.SetABV(fmt.Sprintf("%.1f", abv))
	}
}
