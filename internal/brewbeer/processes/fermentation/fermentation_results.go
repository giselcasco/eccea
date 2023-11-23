package fermentation

import "eccea/internal/brewbeer/characteristics"

type Results struct {
	abv                   float32 // abv alcohol by volume of beer
	flavorCharacteristics []characteristics.FlavorCharacteristic
	smellCharacteristics  []characteristics.SmellCharacteristic
	colorCharacteristics  []characteristics.ColorCharacteristic
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) ABV() float32 {
	return r.abv
}

func (r *Results) SetABV(alcoholByVolume float32) {
	r.abv = alcoholByVolume
}

func (r *Results) AddFlavorCharacteristic(flovorCharactrs characteristics.FlavorCharacteristic) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flovorCharactrs)
}

func (r *Results) AddSmellCharacteristic(smellCharactrs characteristics.SmellCharacteristic) {
	r.smellCharacteristics = append(r.smellCharacteristics, smellCharactrs)
}

func (r *Results) AddColorCharacteristic(colorCharactrs characteristics.ColorCharacteristic) {
	r.colorCharacteristics = append(r.colorCharacteristics, colorCharactrs)
}

func (r *Results) FlavorCharacteristic() []characteristics.FlavorCharacteristic {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []characteristics.SmellCharacteristic {
	return r.smellCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristics.ColorCharacteristic {
	return r.colorCharacteristics
}
