package fermentation

import "eccea/internal/brewbeer/characteristic"

type Results struct {
	abv                   float32 // abv alcohol by volume of beer
	flavorCharacteristics []characteristic.Flavor
	smellCharacteristics  []characteristic.Smell
	colorCharacteristics  []characteristic.Color
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

func (r *Results) AddFlavorCharacteristic(flavor characteristic.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddSmellCharacteristic(smell characteristic.Smell) {
	r.smellCharacteristics = append(r.smellCharacteristics, smell)
}

func (r *Results) AddColorCharacteristic(color characteristic.Color) {
	r.colorCharacteristics = append(r.colorCharacteristics, color)
}

func (r *Results) FlavorCharacteristic() []characteristic.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []characteristic.Smell {
	return r.smellCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristic.Color {
	return r.colorCharacteristics
}
