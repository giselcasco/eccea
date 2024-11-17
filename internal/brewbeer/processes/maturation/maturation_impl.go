package maturation

const detailDescription = "Por el tiempo de maduración se estima que tendrá "

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) EstimateColor(params *Params) string {
	var daysCompare = uint64(7)
	var mapColorIntensity = map[uint64]string{
		7:   "alta turbidez y color opaco.",
		14:  "disminución de la turbidez y color más claro.",
		30:  "alta claridad y color definido.",
		100: "máxima pureza y claridad del color.",
	}

	if params.NumberOfDays > 0 {
		for days := range mapColorIntensity {
			if params.NumberOfDays <= days &&
				(daysCompare > days || daysCompare <= params.NumberOfDays) {
				daysCompare = days
			}
		}
	}
	return detailDescription + mapColorIntensity[daysCompare]
}

func (s *service) EstimateFlavor(params *Params) string {
	var daysCompare = uint64(7)
	var mapFlavorIntensity = map[uint64]string{
		7:   "dejos de mantequilla y mansaza verde.",
		14:  "perfil equilibrado con sensacón de alcohol suave.",
		30:  "textura agradable y sensación de suavidad en el paladar.",
		100: "madurez y armonia de sabores.",
	}

	if params.NumberOfDays > 0 {
		for days := range mapFlavorIntensity {
			if params.NumberOfDays <= days &&
				(daysCompare > days || daysCompare <= params.NumberOfDays) {
				daysCompare = days
			}
		}
	}
	return detailDescription + mapFlavorIntensity[daysCompare]
}

func (s *service) EstimateBeer(params *Params) Results {
	return NewResults(s.EstimateColor(params), s.EstimateFlavor(params))
}
