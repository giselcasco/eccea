package abv

type Estimation struct {
	ABV float32
}

func NewEstimation(abv float32) Estimation {
	return Estimation{
		ABV: abv,
	}
}
