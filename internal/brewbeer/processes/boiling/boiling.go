package boiling

// Service es la interfaz del proceso de coccion
//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output servicemocks  --outpkg servicemocks
type Service interface {
	Do(params *Params, useCaseKey string) (*Results, error)
}
