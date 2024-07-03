package abv

type ABV interface {
	Estimate(params Params) (float64, error)
}
