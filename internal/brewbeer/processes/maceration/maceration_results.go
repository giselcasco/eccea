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

func (r *ColorResults) Color() uint64 {
	return r.colorSRM
}

func (r *ColorResults) SetColor(color uint64) {
	r.colorSRM = color
}

func (r *ColorResults) ColorDescription() string {
	return r.colorDescription
}

func (r *ColorResults) SetColorDescription(colorDes string) {
	r.colorDescription = colorDes
}

func (r *ColorResults) ColorCharacteristic() string {
	return r.colorCharacteristics
}

func (r *ColorResults) SetColorCharacteristic(colorCharacteristic string) {
	r.colorCharacteristics = colorCharacteristic
}
