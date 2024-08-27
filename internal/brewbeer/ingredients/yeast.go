package ingredients

type (
	Yeast struct {
		name              string
		temperatureOfWork TemperatureRange
		timeOfWork        TimeOfWork
	}

	TimeOfWork struct {
		min         float64
		recommended float64
		max         float64
	}

	YeastBuilder struct {
		yeast *Yeast
	}
)

func NewYeastBuilder() *YeastBuilder {
	return &YeastBuilder{
		yeast: &Yeast{},
	}
}

func (yb *YeastBuilder) Name(name string) *YeastBuilder {
	yb.yeast.name = name
	return yb
}

func (yb *YeastBuilder) TemperatureRange(min, max float64) *YeastBuilder {
	yb.yeast.temperatureOfWork = TemperatureRange{
		min: min,
		max: max,
	}
	return yb
}

func (yb *YeastBuilder) TimeOfWork(min, recommended, max float64) *YeastBuilder {
	yb.yeast.timeOfWork = TimeOfWork{
		min:         min,
		recommended: recommended,
		max:         max,
	}
	return yb
}

func (yb *YeastBuilder) Build() *Yeast {
	return yb.yeast
}

// Yeast metodos de la entidad de dominio Levadura
func (yeast *Yeast) Name() string {
	return yeast.name
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
