package color

import "eccea/internal/brewbeer/characteristic"

type Estimation struct {
	ColorSRM             uint16
	ColorCharacteristics []characteristic.Color
}

func NewEstimation(colorSRM uint16, colorCharacteristics []characteristic.Color) Estimation {
	return Estimation{
		ColorSRM:             colorSRM,
		ColorCharacteristics: colorCharacteristics,
	}
}
