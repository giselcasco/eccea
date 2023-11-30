package maceration

type Service interface {
	Do(params *Params) *Results
}
