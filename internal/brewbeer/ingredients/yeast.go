package ingredients

type (
	Yeast struct {
		id                    string
		flavorCharacteristics []Flavor
		smellCharacteristics  []Smell
		colorCharacteristics  []Color
		temperatureOfWork     TemperatureRange
		timeOfWork            TimeOfWork
	}

	TimeOfWork struct {
		min         float64
		recommended float64
		max         float64
	}
)

func (yeast *Yeast) ID() string {
	return yeast.id
}

func (yeast *Yeast) FlavorCharacteristics() []Flavor {
	return yeast.flavorCharacteristics
}

func (yeast *Yeast) SmellCharacteristics() []Smell {
	return yeast.smellCharacteristics
}

func (yeast *Yeast) ColorCharacteristics() []Color {
	return yeast.colorCharacteristics
}

func (yeast *Yeast) TemperatureOfWork() TemperatureRange {
	return yeast.temperatureOfWork
}

func (yeast *Yeast) TimeOfWork() TimeOfWork {
	return yeast.timeOfWork
}

func (time *TimeOfWork) Min() float64 {
	return time.min
}

func (time *TimeOfWork) Max() float64 {
	return time.max
}

func (time *TimeOfWork) Recommended() float64 {
	return time.recommended
}
