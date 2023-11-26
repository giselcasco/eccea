package boiling

import "eccea/internal/brewbeer/characteristics"

type Results struct {
	ibu                      float64 // ibu is the international bitterness unit
	flavorCharacteristics    []characteristics.FlavorCharacteristic
	smellCharacteristics     []characteristics.SmellCharacteristic
	mouthfeelCharacteristics []characteristics.MouthfeelCharacteristic
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) IBU() float64 {
	return r.ibu
}

func (r *Results) SetIBU(ibu float64) {
	r.ibu = ibu
}

func (r *Results) AddFlavorCharacteristic(flovorCharactrs characteristics.FlavorCharacteristic) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flovorCharactrs)
}

func (r *Results) AddSmellCharacteristic(smellCharactrs characteristics.SmellCharacteristic) {
	r.smellCharacteristics = append(r.smellCharacteristics, smellCharactrs)
}

func (r *Results) AddMouthfeelCharacteristic(mouthCharactrs characteristics.MouthfeelCharacteristic) {
	r.mouthfeelCharacteristics = append(r.mouthfeelCharacteristics, mouthCharactrs)
}

func (r *Results) FlavorCharacteristic() []characteristics.FlavorCharacteristic {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []characteristics.SmellCharacteristic {
	return r.smellCharacteristics
}

func (r *Results) MouthfeelCharacteristic() []characteristics.MouthfeelCharacteristic {
	return r.mouthfeelCharacteristics
}
