package maturation

type Service interface {
	EstimateColor(params *Params) string
}
