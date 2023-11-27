package maceration

import (
	"eccea/internal/brewbeer/characteristic"
)

type Results struct {
	colorSRM              uint16
	flavorCharacteristics []characteristic.Flavor
	smellCharacteristics  []characteristic.Smell
	colorCharacteristics  []characteristic.Color
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

func (r *Results) AddFlavorCharacteristic(flavor characteristic.Flavor) {
	r.flavorCharacteristics = append(r.flavorCharacteristics, flavor)
}

func (r *Results) AddColorCharacteristic(color characteristic.Color) {
	r.colorCharacteristics = append(r.colorCharacteristics, color)
}

func (r *Results) AddSmellCharacteristic(smell characteristic.Smell) {
	r.smellCharacteristics = append(r.smellCharacteristics, smell)
}

func (r *Results) FlavorCharacteristic() []characteristic.Flavor {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() []characteristic.Color {
	return r.colorCharacteristics
}

func (r *Results) SmellCharacteristic() []characteristic.Smell {
	return r.smellCharacteristics
}
