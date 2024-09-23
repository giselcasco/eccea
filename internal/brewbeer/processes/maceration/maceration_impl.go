package maceration

import (
	"eccea/internal/brewbeer/ingredients"
	"errors"
	"math"
	"strings"
)

type (
	service struct {
		repo ingredients.Reader
	}

	MaltParam struct {
		NameID          string
		Quantity        float64
		Proportion      float64
		Characteristics []ingredients.Characteristic
		ColorSRM        float64
	}
)

func NewService(repo ingredients.Reader) Service {
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
	return s.buildColorResult(finalColorSRM, malts), nil
}

func (s *service) EstimateFlavor(params *Params) (*FlavorResults, error) {
	if params == nil {
		return nil, errors.New("nil params error")
	}

	malts, err := s.getMalts(params.MaltAdditions, params.TotalQuantity)
	if err != nil {
		return nil, err
	}

	return s.buildFlavorResult(malts), nil
}

func (s *service) getMalts(malts []Malt, totalQuantity float64) ([]MaltParam, error) {
	var (
		proportion float64
		maltParams []MaltParam
	)

	for _, m := range malts {
		malt, err := s.repo.GetMaltByName(m.NameID)
		if err != nil {
			return nil, err
		}

		proportion = (totalQuantity * 100) / m.Quantity
		maltParam := MaltParam{
			NameID:          m.NameID,
			ColorSRM:        malt.ColorSRM(),
			Characteristics: malt.Characteristics(),
			Proportion:      proportion,
			Quantity:        m.Quantity,
		}
		maltParams = append(maltParams, maltParam)
	}
	return maltParams, nil
}

func (s *service) buildFlavorResult(malts []MaltParam) *FlavorResults {
	var (
		flavorDescription string
		smellDescription  string
		result            FlavorResults
	)

	for _, mCharacts := range malts {
		if flavorsDes := s.buildCharacteristicsDescription(mCharacts, "sabor"); len(flavorsDes) > 0 {
			flavorDescription += "La malta " + mCharacts.NameID + " aporta " + flavorsDes
		}
		if smellsDes := s.buildCharacteristicsDescription(mCharacts, "aroma"); len(smellsDes) > 0 {
			smellDescription += "La malta " + mCharacts.NameID + " aporta " + smellsDes
		}
	}

	result.SetFlavorCharacteristic(flavorDescription)
	result.SetSmellCharacteristic(smellDescription)
	return &result
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

func (s *service) buildColorResult(colorFSRM float64, malts []MaltParam) *ColorResults {
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
		if colorDesc := s.buildCharacteristicsDescription(mCharacts, "color"); len(colorDesc) > 0 {
			maltCharacts += "La malta " + mCharacts.NameID + " aporta " + colorDesc

		}
	}

	result.SetColorCharacteristic(maltCharacts)
	return &result
}

func (s *service) buildCharacteristicsDescription(maltParam MaltParam, characteristicType string) string {
	var characteristicsDescription string
	elements := s.countElements(maltParam.Characteristics, characteristicType)
	for index, characterisc := range maltParam.Characteristics {
		if strings.EqualFold(characterisc.CharacteristicType, characteristicType) {
			characteristicsDescription += characterisc.GetDescriptionByProportion(maltParam.Proportion)
			if elements > 1 {
				characteristicsDescription += s.getConnector(index, elements)
			}
		}
	}
	return characteristicsDescription + ".\r"
}

func (s *service) countElements(ccharacts []ingredients.Characteristic, characteristicType string) int {
	amount := 0
	for _, characterisc := range ccharacts {
		if strings.EqualFold(characterisc.CharacteristicType, characteristicType) {
			amount += 1
		}
	}
	return amount
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
