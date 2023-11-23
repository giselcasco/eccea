package maturation

type Maturation interface {
	Do(params *Params) *Results
}
