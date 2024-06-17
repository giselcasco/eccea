package fermentation

// Service es la interfaz del proceso de fermentacion
//
//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output fermentationmocks  --outpkg fermentationmocks
type Service interface {
	Do(params *Params, useCaseKey string) *Results
}
