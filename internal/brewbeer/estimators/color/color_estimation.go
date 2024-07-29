package color

type Estimation struct {
	ColorSRM                  uint64
	ColorDescription          string
	ColorIntensityDescription string
	ColorCharacteristics      string
}

func NewEstimation(colorSRM uint64,
	colorDescription,
	colorIntensityDescription string,
	colorCharacteristics string) *Estimation {
	return &Estimation{
		ColorSRM:                  colorSRM,
		ColorDescription:          colorDescription,
		ColorIntensityDescription: colorIntensityDescription,
		ColorCharacteristics:      colorCharacteristics,
	}
}
