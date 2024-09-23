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

func (r *Results) AddCharacteristic(ccharact ingredients.Characteristic) {
	r.characteristics = append(r.characteristics, ccharact)
}

func (r *Results) Characteristic() []ingredients.Characteristic {
	return r.characteristics
}
