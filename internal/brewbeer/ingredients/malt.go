package ingredients

import (
	"eccea/internal/brewbeer/characteristics"
)

type (
	Malt struct {
		id                    string
		flavorCharacteristics []characteristics.FlavorCharacteristic
		smellCharacteristics  []characteristics.SmellCharacteristic
		colorCharacteristics  []characteristics.ColorCharacteristic
		colorSRM              uint16
		temperatureOfUse      TemperatureRangeOfUse
		extractFineGrind      float32
		extractCoarseGrind    float32
		diastaticPower        float32
	}

	TemperatureRangeOfUse struct {
		min float32
		max float32
	}
)

func (malt *Malt) ID() string {
	return malt.id
}

func (malt *Malt) FlavorCharacteristics() []characteristics.FlavorCharacteristic {
	return malt.flavorCharacteristics
}

func (malt *Malt) SmellCharacteristics() []characteristics.SmellCharacteristic {
	return malt.smellCharacteristics
}

func (malt *Malt) ColorCharacteristics() []characteristics.ColorCharacteristic {
	return malt.colorCharacteristics
}

func (malt *Malt) TemperatureOfUse() TemperatureRangeOfUse {
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

func (temperature *TemperatureRangeOfUse) Min() float32 {
	return temperature.min
}

func (temperature *TemperatureRangeOfUse) Max() float32 {
	return temperature.max
}
