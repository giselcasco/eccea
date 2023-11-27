package flavor

import "eccea/internal/brewbeer/characteristic"

type Estimation struct {
	FlavorCharacteristics    []characteristic.Flavor
	SmellCharacteristics     []characteristic.Smell
	MouthfeelCharacteristics []characteristic.Mouthfeel
}

func NewEstimation(flavorCharacteristics []characteristic.Flavor,
	smellCharacteristics []characteristic.Smell,
	mouthfeelCharacteristics []characteristic.Mouthfeel) Estimation {
	return Estimation{
		FlavorCharacteristics:    flavorCharacteristics,
		SmellCharacteristics:     smellCharacteristics,
		MouthfeelCharacteristics: mouthfeelCharacteristics,
	}
}
