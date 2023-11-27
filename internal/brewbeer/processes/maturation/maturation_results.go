package maturation

import "eccea/internal/brewbeer/characteristic"

type Results struct {
	flavorCharacteristics []characteristic.Flavor
	colorCharacteristics  []characteristic.Color
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) AddFlavor(flavor characteristic.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddColor(color characteristic.Color) {
	r.colorCharacteristics = append(r.colorCharacteristics, color)
}

func (r *Results) FlavorCharacteristic() []characteristic.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristic.Color {
	return r.colorCharacteristics
}
