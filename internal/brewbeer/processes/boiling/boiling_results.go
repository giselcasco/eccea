package boiling

import "eccea/internal/brewbeer/ingredients"

type Results struct {
	ibu                       float64 // ibu is the international bitterness unit
	flavorCharacteristics     []ingredients.Flavor
	smellCharacteristics      []ingredients.Smell
	afterTasteCharacteristics []ingredients.AfterTaste
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

func (r *Results) AddFlavorCharacteristic(flavor ingredients.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddSmellCharacteristic(smell ingredients.Smell) {
	r.smellCharacteristics = append(r.smellCharacteristics, smell)
}

func (r *Results) AddAfterTasteCharacteristic(afterTaste ingredients.AfterTaste) {
	r.afterTasteCharacteristics = append(r.afterTasteCharacteristics, afterTaste)
}

func (r *Results) FlavorCharacteristic() []ingredients.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() []ingredients.Smell {
	return r.smellCharacteristics
}

func (r *Results) AfterTasteCharacteristic() []ingredients.AfterTaste {
	return r.afterTasteCharacteristics
}
