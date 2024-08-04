package maceration

import (
	"eccea/internal/brewbeer/characteristic"
	"eccea/internal/brewbeer/ingredients"
	"errors"
	"math"
)

type (
	service struct {
		repo ingredients.Repository
	}

	MaltParam struct {
		NameID               string
		Quantity             float64
		Proportion           float64
		ColorCharacteristics []characteristic.Color
		ColorSRM             float64
	}
)

func NewService(repo ingredients.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) EstimateColor(params *Params) (*ColorResults, error) {
	if params == nil {
		return nil, errors.New("nil params error")
	}

	malts, err := s.getMalts(params.MaltAdditions, params.TotalQuantity)
	if err != nil {
		return nil, err
	}

	finalColorSRM := s.calculateFinalColorSRM(params, malts)
	return s.buildResult(finalColorSRM, malts), nil
}

func (s *service) getMalts(malts []Malt, totalQuantity float64) ([]MaltParam, error) {
	var (
		proportion float64
		maltParams []MaltParam
	)

	for _, m := range malts {
		malt, err := s.repo.GetMalt(m.NameID)
		if err != nil {
			return nil, err
		}

		proportion = (totalQuantity * 100) / m.Quantity
		maltParam := MaltParam{
			NameID:               m.NameID,
			ColorSRM:             malt.ColorSRM(),
			ColorCharacteristics: malt.ColorCharacteristics(),
			Proportion:           proportion,
			Quantity:             m.Quantity,
		}
		maltParams = append(maltParams, maltParam)
	}
	return maltParams, nil
}

func (s *service) calculateFinalColorSRM(params *Params, malts []MaltParam) float64 {
	var sum float64
	for _, m := range malts {
		sum += (m.Quantity / 1000) * m.ColorSRM
	}

	if params.WortAmount > 0 {
		color := sum / (params.WortAmount * 0.96)
		return color
	}
	return 0
}

func (s *service) buildResult(colorFSRM float64, malts []MaltParam) *ColorResults {
	var (
		colorF       = uint64(math.Round(colorFSRM))
		maltCharacts string
		result       ColorResults
	)
	if colorFSRM > 0 {
		result.SetColor(colorF)
		result.SetColorDescription(s.getDescriptionColor(colorF))
	}

	for _, mCharacts := range malts {
		if len(mCharacts.ColorCharacteristics) > 0 {
			maltCharacts += "La malta " + mCharacts.NameID + " aporta " +
				s.buildColorCharacteristicsDescription(mCharacts)
		}
	}

	result.SetColorCharacteristic(maltCharacts)
	return &result
}
func (s *service) buildColorCharacteristicsDescription(maltParam MaltParam) string {
	var colorCharacteristicsDescription string
	elements := len(maltParam.ColorCharacteristics)
	for index, colorCharacterisc := range maltParam.ColorCharacteristics {
		colorCharacteristicsDescription += colorCharacterisc.GetDescriptionMaltColor(maltParam.Proportion)
		if elements > 1 {
			colorCharacteristicsDescription += s.getConnector(index, elements)
		}
	}
	return colorCharacteristicsDescription
}

func (s *service) getConnector(index, elements int) string {
	if index+2 == elements {
		return " y "
	}
	return ", "
}

func (s *service) getDescriptionColor(colorSRM uint64) string {
	var daysCompare = uint64(3)
	var mapColorDescription = map[uint64]string{
		3:   "Pajoso",
		4:   "Amarillo",
		6:   "Dorado",
		9:   "Ambar",
		14:  "Naranja",
		18:  "Cobrizo",
		22:  "Marrón",
		30:  "Marrón Oscuro",
		100: "Negro",
	}

	for days := range mapColorDescription {
		if colorSRM <= days && (daysCompare > days || daysCompare <= colorSRM) {
			daysCompare = days
		}
	}

	return mapColorDescription[daysCompare]
}
