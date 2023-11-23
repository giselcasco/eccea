package maturation

import "eccea/internal/brewbeer/characteristics"

type Results struct {
	flavorCharacteristics []characteristics.FlavorCharacteristic
	colorCharacteristics  []characteristics.ColorCharacteristic
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) AddFlavorCharacteristic(flovorCharactrs characteristics.FlavorCharacteristic) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flovorCharactrs)
}

func (r *Results) AddColorCharacteristic(colorCharactrs characteristics.ColorCharacteristic) {
	r.colorCharacteristics = append(r.colorCharacteristics, colorCharactrs)
}

func (r *Results) FlavorCharacteristic() []characteristics.FlavorCharacteristic {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristics.ColorCharacteristic {
	return r.colorCharacteristics
}
