package beer

import "eccea/internal/brewbeer/ingredients"

type Estimation struct {
	FlavorCharacteristics     []ingredients.Flavor
	SmellCharacteristics      []ingredients.Smell
	AfterTasteCharacteristics []ingredients.AfterTaste
	ColorCharacteristics      []ingredients.Color
	ColorSRM                  uint16
	ABV                       float32
	IBU                       float32
}

func NewEstimation(flavorCharacteristics []ingredients.Flavor,
	smellCharacteristics []ingredients.Smell,
	afterTasteCharacteristics []ingredients.AfterTaste,
	colorCharacteristics []ingredients.Color,
	colorSRM uint16,
	abv float32,
	ibu float32) Estimation {
	return Estimation{
		FlavorCharacteristics:     flavorCharacteristics,
		SmellCharacteristics:      smellCharacteristics,
		AfterTasteCharacteristics: afterTasteCharacteristics,
		ColorCharacteristics:      colorCharacteristics,
		ColorSRM:                  colorSRM,
		ABV:                       abv,
		IBU:                       ibu,
	}
}
