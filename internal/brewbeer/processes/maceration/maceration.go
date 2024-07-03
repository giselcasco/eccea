package maceration

type Service interface {
	EstimateColor(params *Params) *ColorResults
}
