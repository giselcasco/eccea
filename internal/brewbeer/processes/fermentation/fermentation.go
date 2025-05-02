package fermentation

// Service es la interfaz del proceso de fermentacion

//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output fermentationmocks  --outpkg fermentationmocks
type Service interface {
	CalculateAlcoholByVolume(initialDensity, finalDensity float64) (float64, error)
	EstimateBeer(params *Params) (*Results, error)
}
