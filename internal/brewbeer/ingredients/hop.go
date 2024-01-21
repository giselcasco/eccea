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

func (hop *Hop) SetID(id string) {
	hop.id = id
}

func (hop *Hop) FlavorCharacteristics() []characteristic.Flavor {
	return hop.flavorCharacteristics
}

func (hop *Hop) SetFlavorCharacteristics(flavors []characteristic.Flavor) {
	hop.flavorCharacteristics = flavors
}

func (hop *Hop) SmellCharacteristics() []characteristic.Smell {
	return hop.smellCharacteristics
}

func (hop *Hop) SetSmellCharacteristics(smells []characteristic.Smell) {
	hop.smellCharacteristics = smells
}

func (hop *Hop) MouthfeelCharacteristics() []characteristic.Mouthfeel {
	return hop.mouthfeelCharacteristics
}

func (hop *Hop) SetMouthfeelCharacteristics(mouthfeel []characteristic.Mouthfeel) {
	hop.mouthfeelCharacteristics = mouthfeel
}

func (hop *Hop) BetaAcids() float32 {
	return hop.betaAcids
}

func (hop *Hop) SetBetaAcids(betaAcids float32) {
	hop.betaAcids = betaAcids
}
func (hop *Hop) AlphaAcids() float32 {
	return hop.alphaAcids
}

func (hop *Hop) SetAlphaAcids(alphaAcids float32) {
	hop.alphaAcids = alphaAcids
}
