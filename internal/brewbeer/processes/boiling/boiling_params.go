package boiling

type (
	Params struct {
		totalTime  uint32         // totalTime is the total time in minutes of boiling process.
		wortAmount uint32         // wortAmount is the quantity of wort at the beginning of the process.
		additions  []HopAdditions // additions list of hops included in the process.
	}

	HopAdditions struct {
		id            string
		quantity      float32 // quantity in grams.
		timeFromStart uint32  // timeFromStart in minutes in which the hops enter; for example 30, if you entered 30 minutes from the start of the cooking process.
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

func NewHopAdditions(id string, quantity float32, timeFromStart uint32) *HopAdditions {
	return &HopAdditions{
		id:            id,
		quantity:      quantity,
		timeFromStart: timeFromStart,
	}
}
