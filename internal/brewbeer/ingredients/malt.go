package ingredients

import (
	"eccea/internal/brewbeer/characteristic"
)

type (
	Malt struct {
		nameID                string
		flavorCharacteristics []characteristic.Flavor
		smellCharacteristics  []characteristic.Smell
		colorCharacteristics  []characteristic.Color
		colorSRM              float64
		temperatureOfUse      TemperatureRange
		extractFineGrind      float32
		extractCoarseGrind    float32
		diastaticPower        float32
	}

	TemperatureRange struct {
		min float32
		max float32
	}

	MaltBuilder struct {
		malt *Malt
	}
)

func NewMaltBuilder() *MaltBuilder {
	return &MaltBuilder{
		malt: &Malt{},
	}
}

func (mb *MaltBuilder) NameID(nameID string) {
	mb.malt.nameID = nameID
}

func (mb *MaltBuilder) ColorCharacteristics(ccharact []characteristic.Color) {
	mb.malt.colorCharacteristics = ccharact
}

func (mb *MaltBuilder) ColorSRM(color float64) {
	mb.malt.colorSRM = color
}

func (mb *MaltBuilder) Build() *Malt {
	return mb.malt
}

func (malt *Malt) NameID() string {
	return malt.nameID
}

func (malt *Malt) FlavorCharacteristics() []characteristic.Flavor {
	return malt.flavorCharacteristics
}

func (malt *Malt) SmellCharacteristics() []characteristic.Smell {
	return malt.smellCharacteristics
}

func (malt *Malt) ColorCharacteristics() []characteristic.Color {
	return malt.colorCharacteristics
}

func (malt *Malt) TemperatureOfUse() TemperatureRange {
	return malt.temperatureOfUse
}

func (malt *Malt) ExtractFineGrind() float32 {
	return malt.extractFineGrind
}

func (malt *Malt) ExtractCoarseGrind() float32 {
	return malt.extractCoarseGrind
}

func (malt *Malt) DiastaticPower() float32 {
	return malt.diastaticPower
}

func (malt *Malt) ColorSRM() float64 {
	return malt.colorSRM
}

func (temperature *TemperatureRange) Min() float32 {
	return temperature.min
}

func (temperature *TemperatureRange) Max() float32 {
	return temperature.max
}
