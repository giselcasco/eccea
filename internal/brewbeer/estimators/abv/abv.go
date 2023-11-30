package abv

type ABV interface {
	Estimate(params Params) (Estimation, error)
}
