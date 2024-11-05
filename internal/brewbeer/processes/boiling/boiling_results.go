package boiling

type (
	FlavorResults struct {
		flavorCharacteristics     string
		smellCharacteristics      string
		afterTasteCharacteristics string
	}

	Results struct {
		ibu                       float64
		flavorCharacteristics     string
		smellCharacteristics      string
		afterTasteCharacteristics string
	}
)

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

func (f *Results) SetFlavorCharacteristic(flavor string) {
	f.flavorCharacteristics = flavor
}

func (f *Results) SetSmellCharacteristic(smell string) {
	f.smellCharacteristics = smell
}

func (f *Results) SetAfterTasteCharacteristic(afterTaste string) {
	f.afterTasteCharacteristics = afterTaste
}

func (f *Results) FlavorCharacteristic() string {
	return f.flavorCharacteristics
}

func (f *Results) SmellCCharacteristic() string {
	return f.smellCharacteristics
}

func (f *Results) AfterTasteCharacteristic() string {
	return f.afterTasteCharacteristics
}

func (f *Results) IBU() float64 {
	return f.ibu
}

func (f *Results) SetIBU(ibu float64) {
	f.ibu = ibu
}
