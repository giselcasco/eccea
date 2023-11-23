package maturation

type Params struct {
	totalTime uint32 // totalTime is the total time in hours of maturation process.
}

func NewParams(totalTime uint32) *Params {
	return &Params{
		totalTime: totalTime,
	}
}

func (p Params) TotalTime() uint32 {
	return p.totalTime
}
