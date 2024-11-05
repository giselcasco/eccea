package maturation

//go:generate mockery --name=Service --structname=ServiceMock --case underscore --output maturationmocks  --outpkg maturationmocks
type Service interface {
	EstimateColor(params *Params) string
	EstimateFlavor(params *Params) string
	EstimateBeer(params *Params) Results
}
