package flavor

type Flavor interface {
	Estimate(params Params) (*Estimation, error)
}
