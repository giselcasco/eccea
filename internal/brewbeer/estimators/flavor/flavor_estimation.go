package flavor

type Estimation struct {
	FlavorMaturation          string
	FlavorCharacteristics     string
	SmellCharacteristics      string
	AfterTasteCharacteristics string
}

func NewEstimation(flavorMaturation string, flavorCharacteristics string,
	smellCharacteristics string,
	afterTasteCharacteristics string) *Estimation {
	return &Estimation{
		FlavorMaturation:          flavorMaturation,
		FlavorCharacteristics:     flavorCharacteristics,
		SmellCharacteristics:      smellCharacteristics,
		AfterTasteCharacteristics: afterTasteCharacteristics,
	}
}
