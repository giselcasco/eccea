package fermentation

type service struct {
}

func NewService() Service {
	return &service{
	}
}

/*
    Do es la implementaciòn del proceso de fermentacion, de este se obtienen características de menor percepción
  y el valor de ABV (alcohol by volume)
    useCaseKey hace referencia al caso de uso con el que se consulta al proceso,
 puede ser "abv", "smell", flavor", "color" or "beer"
*/
func (s *service) Do(params *Params) *Results {
    var estimateResults Results
    if abv := calculateIBU(params); abv > 0 {
    	estimateResults.SetABV(fmt.Sprintf("%.1f", abv))
    }
	return estimateResults
}

// estimateABV es el metodo que calcula el abv a partir de los valores en los parametros
func calculateIBU(params *Params, result *Results) {
	return (params.InitialDensity - params.FinalDensity) * 131.25
}