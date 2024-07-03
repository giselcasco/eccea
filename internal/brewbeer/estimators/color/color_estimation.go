package color

import "eccea/internal/brewbeer/characteristic"

type Estimation struct {
	ColorSRM             uint64
	ColorCharacteristics []characteristic.Color
}

func NewEstimation(colorSRM uint64, colorCharacteristics []characteristic.Color) *Estimation {
	return &Estimation{
		ColorSRM:             colorSRM,
		ColorCharacteristics: colorCharacteristics,
	}
}
