package boiling

//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output servicemocks  --outpkg servicemocks
type Service interface {
	Do(params *Params, useCaseKey string) (*Results, error)
}
