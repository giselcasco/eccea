package boiling

// Service es la interfaz del proceso de coccion
//
//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output boilingmocks  --outpkg boilingmocks
type Service interface {
	EstimateIBU(params *Params) (float64, error)
	EstimateFlavor(params *Params) (*FlavorResults, error)
}
