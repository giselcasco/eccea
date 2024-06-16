package boiling

// Service es la interfaz del proceso de coccion
//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output boilingmocks  --outpkg boilingmocks
type Service interface {
	Do(params *Params, useCaseKey string) (*Results, error)
}
