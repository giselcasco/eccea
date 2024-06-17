package abv

type Estimation struct {
	ABV string
}

func NewEstimation(abv string) Estimation {
	return Estimation{
		ABV: abv,
	}
}
