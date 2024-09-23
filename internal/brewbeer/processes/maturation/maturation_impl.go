package maturation

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
	return mapColorIntensity[daysCompare]
}

func (s *service) EstimateFlavor(params *Params) string {
	var daysCompare = uint64(7)
	// TODO modificar mapa, cargar caracteristicas por maduración
	var mapFlavorIntensity = map[uint64]string{
		7:   "alta turbidez y color opaco.",
		14:  "disminución de la turbidez y color más claro.",
		30:  "alta claridad y color definido.",
		100: "máxima pureza y claridad del color.",
	}

	if params.NumberOfDays > 0 {
		for days := range mapFlavorIntensity {
			if params.NumberOfDays <= days &&
				(daysCompare > days || daysCompare <= params.NumberOfDays) {
				daysCompare = days
			}
		}
	}
	return mapFlavorIntensity[daysCompare]
}
