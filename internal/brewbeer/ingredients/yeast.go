package ingredients

import (
	"eccea/internal/brewbeer/characteristic"
)

type (
	Yeast struct {
		id                    string
		flavorCharacteristics []characteristic.Flavor
		smellCharacteristics  []characteristic.Smell
		colorCharacteristics  []characteristic.Color
		temperatureOfWork     TemperatureRange
		timeOfWork            TimeOfWork
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

func (yeast *Yeast) FlavorCharacteristics() []characteristic.Flavor {
	return yeast.flavorCharacteristics
}

func (yeast *Yeast) SmellCharacteristics() []characteristic.Smell {
	return yeast.smellCharacteristics
}

func (yeast *Yeast) ColorCharacteristics() []characteristic.Color {
	return yeast.colorCharacteristics
}

func (yeast *Yeast) TemperatureOfWork() TemperatureRange {
	return yeast.temperatureOfWork
}

func (yeast *Yeast) TimeOfWork() TimeOfWork {
	return yeast.timeOfWork
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
