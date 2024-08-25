package ingredients

type Hop struct {
	name                     string
	flavorCharacteristics    []Flavor
	smellCharacteristics     []Smell
	mouthfeelCharacteristics []Mouthfeel
	alphaAcids               float64
	betaAcids                float64
}

func (hop *Hop) Name() string {
	return hop.name
}

func (hop *Hop) SetID(name string) {
	hop.name = name
}

func (hop *Hop) FlavorCharacteristics() []Flavor {
	return hop.flavorCharacteristics
}

func (hop *Hop) SetFlavorCharacteristics(flavors []Flavor) {
	hop.flavorCharacteristics = flavors
}

func (hop *Hop) SmellCharacteristics() []Smell {
	return hop.smellCharacteristics
}

func (hop *Hop) SetSmellCharacteristics(smells []Smell) {
	hop.smellCharacteristics = smells
}

func (hop *Hop) MouthfeelCharacteristics() []Mouthfeel {
	return hop.mouthfeelCharacteristics
}

func (hop *Hop) SetMouthfeelCharacteristics(mouthfeel []Mouthfeel) {
	hop.mouthfeelCharacteristics = mouthfeel
}

func (hop *Hop) BetaAcids() float64 {
	return hop.betaAcids
}

func (hop *Hop) SetBetaAcids(betaAcids float64) {
	hop.betaAcids = betaAcids
}
func (hop *Hop) AlphaAcids() float64 {
	return hop.alphaAcids
}

func (hop *Hop) SetAlphaAcids(alphaAcids float64) {
	hop.alphaAcids = alphaAcids
}
