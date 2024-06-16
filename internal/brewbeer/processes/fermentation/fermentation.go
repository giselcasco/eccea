package fermentation

// Service es la interfaz del proceso de fermentacion
//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output fermentationocks  --outpkg fermentationocks
type Service interface {
	Do(params *Params) *Results
}
