package ibu

type Estimation struct {
	IBU float32
}

func NewEstimation(ibu float32) Estimation {
	return Estimation{
		IBU: ibu,
	}
}
