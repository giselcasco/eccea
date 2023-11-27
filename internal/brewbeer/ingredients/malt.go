package ingredients

import (
	"eccea/internal/brewbeer/characteristic"
)

type (
	Malt struct {
		id                    string
		flavorCharacteristics []characteristic.Flavor
		smellCharacteristics  []characteristic.Smell
		colorCharacteristics  []characteristic.Color
		colorSRM              uint16
		temperatureOfUse      TemperatureRange
		extractFineGrind      float32
		extractCoarseGrind    float32
		diastaticPower        float32
	}

	TemperatureRange struct {
		min float32
		max float32
	}
)

func (malt *Malt) ID() string {
	return malt.id
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

func (temperature *TemperatureRange) Min() float32 {
	return temperature.min
}

func (temperature *TemperatureRange) Max() float32 {
	return temperature.max
}
