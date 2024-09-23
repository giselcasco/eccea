package maceration

type (
	Results struct {
		colorSRM              uint64
		flavorCharacteristics string
		smellCharacteristics  string
		colorCharacteristics  string
	}

	ColorResults struct {
		colorSRM             uint64
		colorDescription     string
		colorCharacteristics string
	}

	FlavorResults struct {
		flavorCharacteristics string
		smellCharacteristics  string
	}
)

func NewResults() *Results {
	return &Results{}
}

func (r *Results) Color() uint64 {
	return r.colorSRM
}

func (r *Results) SetColor(color uint64) {
	r.colorSRM = color
}

func (r *Results) SetFlavorCharacteristic(flavorCharacteristic string) {
	r.flavorCharacteristics = flavorCharacteristic
}

func (r *Results) SetColorCharacteristic(colorCharacteristic string) {
	r.colorCharacteristics = colorCharacteristic
}

func (r *Results) SetSmellCharacteristic(smellCharacteristic string) {
	r.smellCharacteristics = smellCharacteristic
}

func (r *Results) FlavorCharacteristic() string {
	return r.flavorCharacteristics
}

func (r *Results) ColorCharacteristic() string {
	return r.colorCharacteristics
}

func (r *Results) SmellCharacteristic() string {
	return r.smellCharacteristics
}

func NewColorResults() *ColorResults {
	return &ColorResults{}
}

func (c *ColorResults) Color() uint64 {
	return c.colorSRM
}

func (c *ColorResults) SetColor(color uint64) {
	c.colorSRM = color
}

func (c *ColorResults) ColorDescription() string {
	return c.colorDescription
}

func (c *ColorResults) SetColorDescription(colorDes string) {
	c.colorDescription = colorDes
}

func (c *ColorResults) ColorCharacteristic() string {
	return c.colorCharacteristics
}

func (c *ColorResults) SetColorCharacteristic(colorCharacteristic string) {
	c.colorCharacteristics = colorCharacteristic
}

func NewFlavorResults() *FlavorResults {
	return &FlavorResults{}
}

func (f *FlavorResults) SetFlavorCharacteristic(flavorCharacteristic string) {
	f.flavorCharacteristics = flavorCharacteristic
}

func (f *FlavorResults) SetSmellCharacteristic(smellCharacteristics string) {
	f.smellCharacteristics = smellCharacteristics
}

func (f *FlavorResults) FlavorCharacteristic() string {
	return f.flavorCharacteristics
}

func (f *FlavorResults) SmellCharacteristic() string {
	return f.smellCharacteristics
}
