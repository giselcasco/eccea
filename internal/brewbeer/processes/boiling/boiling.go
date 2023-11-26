package boiling

type Boiling interface {
	Do(params *Params) (*Results, error)
}
