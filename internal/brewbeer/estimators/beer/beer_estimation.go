package beer

import "eccea/internal/brewbeer/characteristic"

type Estimation struct {
	FlavorCharacteristics    []characteristic.Flavor
	SmellCharacteristics     []characteristic.Smell
	MouthfeelCharacteristics []characteristic.Mouthfeel
	ColorCharacteristics     []characteristic.Color
	ColorSRM                 uint16
	ABV                      float32
	IBU                      float32
}

func NewEstimation(flavorCharacteristics []characteristic.Flavor,
	smellCharacteristics []characteristic.Smell,
	mouthfeelCharacteristics []characteristic.Mouthfeel,
	colorCharacteristics []characteristic.Color,
	colorSRM uint16,
	abv float32,
	ibu float32) Estimation {
	return Estimation{
		FlavorCharacteristics:    flavorCharacteristics,
		SmellCharacteristics:     smellCharacteristics,
		MouthfeelCharacteristics: mouthfeelCharacteristics,
		ColorCharacteristics:     colorCharacteristics,
		ColorSRM:                 colorSRM,
		ABV:                      abv,
		IBU:                      ibu,
	}
}
