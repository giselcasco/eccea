package beer

import "eccea/internal/brewbeer/characteristics"

type Estimation struct {
	FlavorCharacteristics    []characteristics.FlavorCharacteristic
	SmellCharacteristics     []characteristics.SmellCharacteristic
	MouthfeelCharacteristics []characteristics.MouthfeelCharacteristic
	ColorCharacteristics     []characteristics.ColorCharacteristic
	ColorSRM                 uint16
	ABV                      float32
	IBU                      float32
}

func NewEstimation(flavorCharacteristics []characteristics.FlavorCharacteristic,
	smellCharacteristics []characteristics.SmellCharacteristic,
	mouthfeelCharacteristics []characteristics.MouthfeelCharacteristic,
	colorCharacteristics []characteristics.ColorCharacteristic,
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
