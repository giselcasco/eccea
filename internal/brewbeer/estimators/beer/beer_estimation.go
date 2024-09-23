package beer

import "eccea/internal/brewbeer/ingredients"

type Estimation struct {
	Characteristics []ingredients.Characteristic
	ColorSRM        uint16
	ABV             float32
	IBU             float32
}

func NewEstimation(
	characteristics []ingredients.Characteristic,
	colorSRM uint16,
	abv float32,
	ibu float32) Estimation {
	return Estimation{
		Characteristics: characteristics,
		ColorSRM:        colorSRM,
		ABV:             abv,
		IBU:             ibu,
	}
}
