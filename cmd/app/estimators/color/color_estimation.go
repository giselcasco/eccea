package color

import "eccea/internal/brewbeer/characteristics"

type Estimation struct {
	ColorSRM             uint16
	ColorCharacteristics []characteristics.ColorCharacteristic
}

func NewEstimation(colorSRM uint16, colorCharacteristics []characteristics.ColorCharacteristic) Estimation {
	return Estimation{
		ColorSRM:             colorSRM,
		ColorCharacteristics: colorCharacteristics,
	}
}
