package maturation

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) EstimateColor(params *Params) string {
	var daysCompare = uint64(7)
	var mapColorIntensity = map[uint64]string{
		7:   "Alta turbidez, color opaco.",
		14:  "Disminución de la turbidez, color más claro.",
		30:  "Alta claridad, color definido.",
		100: "Máxima pureza y claridad del color.",
	}

	if params.NumberOfDays > 0 {
		for days := range mapColorIntensity {
			if params.NumberOfDays <= days &&
				(daysCompare > days || daysCompare <= params.NumberOfDays) {
				daysCompare = days
			}
		}
	}
	return mapColorIntensity[daysCompare]
}
