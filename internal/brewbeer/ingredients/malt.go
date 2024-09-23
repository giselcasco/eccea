package ingredients

type (
	Malt struct {
		name               string
		characteristics    []Characteristic
		colorSRM           float64
		temperatureOfUse   TemperatureRange
		extractFineGrind   float64
		extractCoarseGrind float64
		diastaticPower     float64
	}

	TemperatureRange struct {
		min float64
		max float64
	}

	MaltBuilder struct {
		malt *Malt
	}
)

// MaltBuilder constructor de la entidad malta
func NewMaltBuilder() *MaltBuilder {
	return &MaltBuilder{
		malt: &Malt{},
	}
}

func (mb *MaltBuilder) Name(name string) *MaltBuilder {
	mb.malt.name = name
	return mb
}

func (mb *MaltBuilder) Characteristics(ccharact []Characteristic) *MaltBuilder {
	mb.malt.characteristics = ccharact
	return mb
}

func (mb *MaltBuilder) ColorSRM(color float64) *MaltBuilder {
	mb.malt.colorSRM = color
	return mb
}

func (mb *MaltBuilder) DiastaticPower(diastaticPower float64) *MaltBuilder {
	mb.malt.diastaticPower = diastaticPower
	return mb
}

func (mb *MaltBuilder) ExtractCoarseGrind(extractCoarseGrind float64) *MaltBuilder {
	mb.malt.extractCoarseGrind = extractCoarseGrind
	return mb
}

func (mb *MaltBuilder) ExtractFineGrind(extractFineGrind float64) *MaltBuilder {
	mb.malt.extractFineGrind = extractFineGrind
	return mb
}

func (mb *MaltBuilder) TemperatureRange(min, max float64) *MaltBuilder {
	mb.malt.temperatureOfUse = TemperatureRange{
		min: min,
		max: max,
	}
	return mb
}

func (mb *MaltBuilder) Build() *Malt {
	return mb.malt
}

// Malt metodos de la entidad de dominio
func (malt *Malt) Name() string {
	return malt.name
}

func (malt *Malt) Characteristics() []Characteristic {
	return malt.characteristics
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
