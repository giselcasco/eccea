package ingredients

import (
	"eccea/internal/brewbeer/characteristic"
)

type Hop struct {
	id                       string
	flavorCharacteristics    []characteristic.Flavor
	smellCharacteristics     []characteristic.Smell
	mouthfeelCharacteristics []characteristic.Mouthfeel
	alphaAcids               float32
	betaAcids                float32
}

func (hop *Hop) ID() string {
	return hop.id
}

func (hop *Hop) FlavorCharacteristics() []characteristic.Flavor {
	return hop.flavorCharacteristics
}

func (hop *Hop) SmellCharacteristics() []characteristic.Smell {
	return hop.smellCharacteristics
}

func (hop *Hop) MouthfeelCharacteristics() []characteristic.Mouthfeel {
	return hop.mouthfeelCharacteristics
}

func (hop *Hop) BetaAcids() float32 {
	return hop.betaAcids
}

func (hop *Hop) AlphaAcids() float32 {
	return hop.alphaAcids
}
