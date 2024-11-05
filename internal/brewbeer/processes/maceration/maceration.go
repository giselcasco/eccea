package maceration

//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output macerationmocks  --outpkg macerationmocks
type Service interface {
	EstimateColor(params *Params) (*ColorResults, error)
	EstimateFlavor(params *Params) (*FlavorResults, error)
	EstimateBeer(params *Params) (*Results, error)
}
