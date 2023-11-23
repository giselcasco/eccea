package maceration

type Maceration interface {
	Do(params *Params) *Results
}
