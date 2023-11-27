package beer

type Beer interface {
	Estimate(params Params) Estimation
}
