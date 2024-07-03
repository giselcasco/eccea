package color

type Color interface {
	Estimate(params Params) *Estimation
}
