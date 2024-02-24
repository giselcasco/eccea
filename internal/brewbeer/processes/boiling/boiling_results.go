package boiling

import "eccea/internal/brewbeer/characteristic"

type Results struct {
	ibu                      string // ibu is the international bitterness unit
	flavorCharacteristics    []characteristic.Flavor
	smellCharacteristics     []characteristic.Smell
	mouthfeelCharacteristics []characteristic.Mouthfeel
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) IBU() string {
	return r.ibu
}

func (r *Results) SetIBU(ibu string) {
	r.ibu = ibu
}

func (r *Results) AddFlavorCharacteristic(flavor characteristic.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddSmellCharacteristic(smell characteristic.Smell) {
	r.smellCharacteristics = append(r.smellCharacteristics, smell)
}

func (r *Results) AddMouthfeelCharacteristic(mouthfeel characteristic.Mouthfeel) {
	r.mouthfeelCharacteristics = append(r.mouthfeelCharacteristics, mouthfeel)
}

func (r *Results) FlavorCharacteristic() []characteristic.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []characteristic.Smell {
	return r.smellCharacteristics
}

func (r *Results) MouthfeelCharacteristic() []characteristic.Mouthfeel {
	return r.mouthfeelCharacteristics
}
