package maturation

type Service interface {
	Do(params *Params) *Results
}
