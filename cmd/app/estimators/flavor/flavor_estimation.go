package flavor

import "eccea/internal/brewbeer/characteristics"

type Estimation struct {
	FlavorCharacteristics    []characteristics.FlavorCharacteristic
	SmellCharacteristics     []characteristics.SmellCharacteristic
	MouthfeelCharacteristics []characteristics.MouthfeelCharacteristic
}

func NewEstimation(flavorCharacteristics []characteristics.FlavorCharacteristic,
	smellCharacteristics []characteristics.SmellCharacteristic,
	mouthfeelCharacteristics []characteristics.MouthfeelCharacteristic) Estimation {
	return Estimation{
		FlavorCharacteristics:    flavorCharacteristics,
		SmellCharacteristics:     smellCharacteristics,
		MouthfeelCharacteristics: mouthfeelCharacteristics,
	}
}
