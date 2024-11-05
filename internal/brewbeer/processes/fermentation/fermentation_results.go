package fermentation

type Results struct {
	abv                   float64 // abv alcohol by volume of beer
	flavorCharacteristics string
	smellCharacteristics  string
	colorCharacteristics  string
}

func (r *Results) ABV() float64 {
	return r.abv
}

func (r *Results) SetABV(alcoholByVolume float64) {
	r.abv = alcoholByVolume
}

func (r *Results) SetFlavorCharacteristic(flavor string) {
	r.flavorCharacteristics = flavor
}

func (r *Results) SetSmellCharacteristic(smell string) {
	r.smellCharacteristics = smell
}

func (r *Results) SetColorCharacteristic(color string) {
	r.colorCharacteristics = color
}

func (r *Results) FlavorCharacteristic() string {
	return r.flavorCharacteristics
}

func (r *Results) SmellCCharacteristic() string {
	return r.smellCharacteristics
}

func (r *Results) ColorCharacteristic() string {
	return r.colorCharacteristics
}
