package ingredients

import (
	"eccea/internal/brewbeer/characteristics"
)

type (
	Yeast struct {
		id                    string
		flavorCharacteristics []characteristics.FlavorCharacteristic
		smellCharacteristics  []characteristics.SmellCharacteristic
		colorCharacteristics  []characteristics.ColorCharacteristic
		temperatureOfWork     TemperatureRangeOfWork
		timeOfWork            TimeOfWork
	}

	TemperatureRangeOfWork struct {
		min float32
		max float32
	}

	TimeOfWork struct {
		min         float32
		recommended float32
		max         float32
	}
)

func (yeast *Yeast) ID() string {
	return yeast.id
}

func (yeast *Yeast) FlavorCharacteristics() []characteristics.FlavorCharacteristic {
	return yeast.flavorCharacteristics
}

func (yeast *Yeast) SmellCharacteristics() []characteristics.SmellCharacteristic {
	return yeast.smellCharacteristics
}

func (yeast *Yeast) ColorCharacteristics() []characteristics.ColorCharacteristic {
	return yeast.colorCharacteristics
}

func (yeast *Yeast) TemperatureOfWork() TemperatureRangeOfWork {
	return yeast.temperatureOfWork
}

func (yeast *Yeast) TimeOfWork() TimeOfWork {
	return yeast.timeOfWork
}

func (temperature *TemperatureRangeOfWork) Min() float32 {
	return temperature.min
}

func (temperature *TemperatureRangeOfWork) Max() float32 {
	return temperature.max
}

func (time *TimeOfWork) Min() float32 {
	return time.min
}

func (time *TimeOfWork) Max() float32 {
	return time.max
}

func (time *TimeOfWork) Recommended() float32 {
	return time.recommended
}
