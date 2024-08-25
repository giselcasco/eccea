package ingredients

type (
	Malt struct {
		name                  string
		flavorCharacteristics []Flavor
		smellCharacteristics  []Smell
		colorCharacteristics  []Color
		colorSRM              float64
		temperatureOfUse      TemperatureRange
		extractFineGrind      float64
		extractCoarseGrind    float64
		diastaticPower        float64
	}

	TemperatureRange struct {
		min float64
		max float64
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

func (mb *MaltBuilder) NameID(name string) {
	mb.malt.name = name
}

func (mb *MaltBuilder) ColorCharacteristics(ccharact []Color) {
	mb.malt.colorCharacteristics = ccharact
}

func (mb *MaltBuilder) ColorSRM(color float64) {
	mb.malt.colorSRM = color
}

func (mb *MaltBuilder) Build() *Malt {
	return mb.malt
}

func (malt *Malt) Name() string {
	return malt.name
}

func (malt *Malt) FlavorCharacteristics() []Flavor {
	return malt.flavorCharacteristics
}

func (malt *Malt) SmellCharacteristics() []Smell {
	return malt.smellCharacteristics
}

func (malt *Malt) ColorCharacteristics() []Color {
	return malt.colorCharacteristics
}

func (malt *Malt) TemperatureOfUse() TemperatureRange {
	return malt.temperatureOfUse
}

func (malt *Malt) ExtractFineGrind() float64 {
	return malt.extractFineGrind
}

func (malt *Malt) ExtractCoarseGrind() float64 {
	return malt.extractCoarseGrind
}

func (malt *Malt) DiastaticPower() float64 {
	return malt.diastaticPower
}

func (malt *Malt) ColorSRM() float64 {
	return malt.colorSRM
}

func (temperature *TemperatureRange) Min() float64 {
	return temperature.min
}

func (temperature *TemperatureRange) Max() float64 {
	return temperature.max
}
