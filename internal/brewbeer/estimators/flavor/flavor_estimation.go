package flavor

type Estimation struct {
	FlavorCharacteristics     string
	SmellCharacteristics      string
	AfterTasteCharacteristics string
}

func NewEstimation(flavorCharacteristics string,
	smellCharacteristics string,
	afterTasteCharacteristics string) *Estimation {
	return &Estimation{
		FlavorCharacteristics:     flavorCharacteristics,
		SmellCharacteristics:      smellCharacteristics,
		AfterTasteCharacteristics: afterTasteCharacteristics,
	}
}
