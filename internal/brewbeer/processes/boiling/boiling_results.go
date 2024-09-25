package boiling

import "eccea/internal/brewbeer/ingredients"

type (
	Results struct {
		ibu             float64 // ibu is the international bitterness unit
		characteristics []ingredients.Characteristic
	}

	FlavorResults struct {
		flavorCharacteristics     string
		smellCharacteristics      string
		afterTasteCharacteristics string
		contributionDescription   string
	}
)

func NewResults() *Results {
	return &Results{}
}

func (r *Results) IBU() float64 {
	return r.ibu
}

func (r *Results) SetIBU(ibu float64) {
	r.ibu = ibu
}

func (f *FlavorResults) SetFlavorCharacteristic(flavor string) {
	f.flavorCharacteristics = flavor
}

func (f *FlavorResults) SetSmellCharacteristic(smell string) {
	f.smellCharacteristics = smell
}

func (f *FlavorResults) SetAfterTasteCharacteristic(afterTaste string) {
	f.afterTasteCharacteristics = afterTaste
}

func (f *FlavorResults) FlavorCharacteristic() string {
	return f.flavorCharacteristics
}

func (f *FlavorResults) SmellCCharacteristic() string {
	return f.smellCharacteristics
}

func (f *FlavorResults) AfterTasteCharacteristic() string {
	return f.afterTasteCharacteristics
}
