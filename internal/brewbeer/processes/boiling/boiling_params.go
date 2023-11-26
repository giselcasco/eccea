package boiling

type (
	Params struct {
		totalTime      uint32 // totalTime is the total time in minutes of boiling process.
		wortAmount     uint32 // wortAmount is the quantity of wort at the beginning of the process.
		initialDensity uint32
		volume         uint32         // volume in liters.
		additions      []HopAdditions // additions list of hops included in the process.
	}

	HopAdditions struct {
		id         string
		quantity   float32 // quantity in grams.
		timeOfWork uint32  // timeOfWork in minutes.
	}
)

func NewParams(totalTime uint32, additions []HopAdditions) *Params {
	return &Params{
		totalTime: totalTime,
		additions: additions,
	}
}

func (p Params) TotalTime() uint32 {
	return p.totalTime
}

func (p Params) Additions() []HopAdditions {
	return p.additions
}

func NewHopAdditions(id string, quantity float32, timeOfWork uint32) *HopAdditions {
	return &HopAdditions{
		id:         id,
		quantity:   quantity,
		timeOfWork: timeOfWork,
	}
}
