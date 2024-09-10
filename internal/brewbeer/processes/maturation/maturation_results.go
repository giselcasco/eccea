package maturation

import "eccea/internal/brewbeer/ingredients"

type (
	Results struct {
		flavorCharacteristics []ingredients.Flavor
		colorCharacteristics  []ingredients.Color
	}

	ColorResults struct {
		colorIntensity string
	}

	FlavorResults struct {
		flavorCharacteristics []ingredients.Flavor
	}
)

func NewResults() *Results {
	return &Results{}
}

func (r *Results) AddFlavor(flavor ingredients.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddColor(color ingredients.Color) {
	r.colorCharacteristics = append(r.colorCharacteristics, color)
}

func (r *Results) FlavorCharacteristic() []ingredients.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() []ingredients.Color {
	return r.colorCharacteristics
}

func (r *ColorResults) ColorIntensity() string {
	return r.colorIntensity
}
