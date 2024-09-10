package fermentation

import "eccea/internal/brewbeer/ingredients"

type Results struct {
	abv                   float64 // abv alcohol by volume of beer
	flavorCharacteristics []ingredients.Flavor
	smellCharacteristics  []ingredients.Smell
	colorCharacteristics  []ingredients.Color
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) ABV() float64 {
	return r.abv
}

func (r *Results) SetABV(alcoholByVolume float64) {
	r.abv = alcoholByVolume
}

func (r *Results) AddFlavorCharacteristic(flavor ingredients.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddSmellCharacteristic(smell ingredients.Smell) {
	r.smellCharacteristics = append(r.smellCharacteristics, smell)
}

func (r *Results) AddColorCharacteristic(color ingredients.Color) {
	r.colorCharacteristics = append(r.colorCharacteristics, color)
}

func (r *Results) FlavorCharacteristic() []ingredients.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []ingredients.Smell {
	return r.smellCharacteristics
}

func (r *Results) ColorCharacteristic() []ingredients.Color {
	return r.colorCharacteristics
}
