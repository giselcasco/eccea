package ingredients

import (
	"eccea/internal/brewbeer/characteristics"
)

type Hop struct {
	id                       string
	flavorCharacteristics    []characteristics.FlavorCharacteristic
	smellCharacteristics     []characteristics.SmellCharacteristic
	mouthfeelCharacteristics []characteristics.MouthfeelCharacteristic
	alphaAcids               float32
	betaAcids                float32
}

func (hop *Hop) ID() string {
	return hop.id
}

func (hop *Hop) FlavorCharacteristics() []characteristics.FlavorCharacteristic {
	return hop.flavorCharacteristics
}

func (hop *Hop) SmellCharacteristics() []characteristics.SmellCharacteristic {
	return hop.smellCharacteristics
}

func (hop *Hop) MouthfeelCharacteristics() []characteristics.MouthfeelCharacteristic {
	return hop.mouthfeelCharacteristics
}

func (hop *Hop) BetaAcids() float32 {
	return hop.betaAcids
}

func (hop *Hop) AlphaAcids() float32 {
	return hop.alphaAcids
}
