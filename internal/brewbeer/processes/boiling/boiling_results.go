package boiling

type FlavorResults struct {
	flavorCharacteristics     string
	smellCharacteristics      string
	afterTasteCharacteristics string
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
