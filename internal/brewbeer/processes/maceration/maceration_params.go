package maceration

type (
	Params struct {
		totalTime uint32 // totalTime is the total in minutes time of maceration process.
		water     float32
		additions []MaltAddition // additions list of malts included in the process.
	}

	MaltAddition struct {
		id         string
		quantity   float32 // quantity in grams.
		timeOfWork uint32  // timeOfWork in minutes.
	}
)

func NewParams(totalTime uint32, water float32, additions []MaltAddition) *Params {
	return &Params{
		totalTime: totalTime,
		water:     water,
		additions: additions,
	}
}

func (p Params) TotalTime() uint32 {
	return p.totalTime
}

func (p Params) Water() float32 {
	return p.water
}

func (p Params) Additions() []MaltAddition {
	return p.additions
}

func NewMaltAddition(id string, quantity float32, timeOfWork uint32) *MaltAddition {
	return &MaltAddition{
		id:         id,
		quantity:   quantity,
		timeOfWork: timeOfWork,
	}
}
