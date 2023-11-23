package fermentation

type Fermentation interface {
	Do(params *Params) *Results
}
