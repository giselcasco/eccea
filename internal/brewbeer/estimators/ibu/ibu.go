package ibu

type IBU interface {
	Estimate(params Params) Estimation
}
