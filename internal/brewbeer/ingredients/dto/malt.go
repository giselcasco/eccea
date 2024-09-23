package dto

import "eccea/internal/brewbeer/ingredients"

type MaltResponse struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	ColorSRM           float64 `json:"color"`
	TemperatureMin     float64 `json:"temp_min"`
	TemperatureMax     float64 `json:"temp_max"`
	ExtractFineGrind   float64 `json:"extract_fine_grind"`
	ExtractCoarseGrind float64 `json:"extract_coarse_grind"`
	// DiastaticPower bajo o nulo en maltas caramelizadas o tostadas
	DiastaticPower  float64                  `json:"diastatic_power"`
	Characteristics []CharacteristicResponse `json:"characteristics"`
}

func (m *MaltResponse) ToDomain() *ingredients.Malt {
	ccharacts := []ingredients.Characteristic{}
	for _, cc := range m.Characteristics {
		ccharact := ingredients.Characteristic{
			TypeAdjetives:      cc.AdjetiveID,
			CharacteristicType: cc.Type,
			Description:        cc.Description,
			Contribution:       cc.Contribution,
		}
		ccharacts = append(ccharacts, ccharact)
	}

	return ingredients.NewMaltBuilder().
		ColorSRM(m.ColorSRM).
		Characteristics(ccharacts).
		TemperatureRange(m.TemperatureMin, m.TemperatureMax).
		Name(m.Name).
		Build()
}
