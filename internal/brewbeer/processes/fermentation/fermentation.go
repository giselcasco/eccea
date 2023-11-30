package fermentation

type Service interface {
	Do(params *Params) *Results
}
