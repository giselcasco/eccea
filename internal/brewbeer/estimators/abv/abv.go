package abv

type ABV interface {
	EstimateAlcoholByVolume(params Params) (float64, error)
}
