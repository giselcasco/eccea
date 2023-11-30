package boiling

type Service interface {
	Do(params *Params, useCaseKey string) (*Results, error)
}
