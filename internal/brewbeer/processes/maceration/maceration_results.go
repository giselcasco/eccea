package maceration

import (
	"eccea/internal/brewbeer/characteristics"
)

type Results struct {
	colorSRM              uint16
	flavorCharacteristics []characteristics.FlavorCharacteristic
	smellCharacteristics  []characteristics.SmellCharacteristic
	colorCharacteristics  []characteristics.ColorCharacteristic
}

func NewResults() *Results {
	return &Results{}
}

func (r *Results) Color() uint16 {
	return r.colorSRM
}

func (r *Results) SetColor(color uint16) {
	r.colorSRM = color
}

func (r *Results) AddFlavorCharacteristic(flovorCharactrs characteristics.FlavorCharacteristic) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flovorCharactrs)
}

func (r *Results) AddColorCharacteristic(colorCharactrs characteristics.ColorCharacteristic) {
	r.colorCharacteristics = append(r.colorCharacteristics, colorCharactrs)
}

func (r *Results) FlavorCharacteristic() []characteristics.FlavorCharacteristic {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristics.ColorCharacteristic {
	return r.colorCharacteristics
}
