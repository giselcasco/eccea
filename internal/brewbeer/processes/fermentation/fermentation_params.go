package fermentation

type (
	Params struct {
		totalTime      uint32 // totalTime is the total time in minutes of fermentation process.
		initialDensity uint32
		finalDensity   uint32
		yeast          YeastParam
	}

	YeastParam struct {
		id         string
		quantity   float32 // quantity in grams.
		timeOfWork uint32  // timeOfWork in minutes.
	}
)

func NewParams(totalTime, initialDensity, finalDensity uint32, yeast YeastParam) *Params {
	return &Params{
		totalTime:      totalTime,
		initialDensity: initialDensity,
		finalDensity:   finalDensity,
		yeast:          yeast,
	}
}

func (p Params) TotalTime() uint32 {
	return p.totalTime
}

func (p Params) InitialDensity() uint32 {
	return p.initialDensity
}

func (p Params) FinalDensity() uint32 {
	return p.finalDensity
}

func (p Params) Yeast() YeastParam {
	return p.yeast
}

func NewYeastParam(id string, quantity float32, timeOfWork uint32) *YeastParam {
	return &YeastParam{
		id:         id,
		quantity:   quantity,
		timeOfWork: timeOfWork,
	}
}
