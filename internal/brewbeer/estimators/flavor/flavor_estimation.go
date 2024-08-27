package flavor

import "eccea/internal/brewbeer/ingredients"

type Estimation struct {
	FlavorCharacteristics     []ingredients.Flavor
	SmellCharacteristics      []ingredients.Smell
	AfterTasteCharacteristics []ingredients.AfterTaste
}

func NewEstimation(flavorCharacteristics []ingredients.Flavor,
	smellCharacteristics []ingredients.Smell,
	afterTasteCharacteristics []ingredients.AfterTaste) Estimation {
	return Estimation{
		FlavorCharacteristics:     flavorCharacteristics,
		SmellCharacteristics:      smellCharacteristics,
		AfterTasteCharacteristics: afterTasteCharacteristics,
	}
}
