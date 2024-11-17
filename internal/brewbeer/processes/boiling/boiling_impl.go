package boiling

import (
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/repository"
	"math"
	"strings"
)

// service es el implementador de los metodos de proceso de coccion
type (
	service struct {
		repo repository.Reader
	}

	HopParams struct {
		NameID          string  // NameID es el nombre con que se conoce al lupulo.
		AlphaAcids      float64 // AlphaAcids que tiene el lúpulo.
		Proportion      float64 // Proporción que representa del total de lúpulos.
		Characteristics []ingredients.Characteristic
		Quantity        float64 // Quantity en gramos.
		TimeOfWork      uint64  // TimeOfWork en minutos.
	}
)

func NewService(repo repository.Reader) Service {
	return &service{
		repo: repo,
	}
}

/*
EstimateIBU es el metodo que calcula el ibu a partir de los valores en los parametros
ingresados por el usuario.
*/
func (s *service) EstimateIBU(params *Params) (float64, error) {
	var ibu float64

	hops, err := s.getHops(params.HopAdditions)
	if err != nil {
		return ibu, err
	}

	for _, hopAddition := range hops {
		ibu += s.calculateIBU(params, hopAddition)
	}

	return ibu, nil
}

/*
EstimateIBU es el metodo que calcula el ibu a partir de los valores en los parametros
ingresados por el usuario.
*/
func (s *service) EstimateFlavor(params *Params) (*FlavorResults, error) {
	hops, herr := s.getHops(params.HopAdditions)
	if herr != nil {
		return nil, herr
	}

	return s.buildFlavorResult(hops), nil
}

func (s *service) EstimateBeer(params *Params) (*Results, error) {
	var ibu float64

	hops, herr := s.getHops(params.HopAdditions)
	if herr != nil {
		return nil, herr
	}

	for _, hopAddition := range hops {
		ibu += s.calculateIBU(params, hopAddition)
	}

	flavorResult := s.buildFlavorResult(hops)
	return s.buildResult(ibu, flavorResult), nil
}

func (s *service) buildFlavorResult(hops []HopParams) *FlavorResults {
	var (
		flavorDescription     string
		smellDescription      string
		afterTasteDescription string
		result                FlavorResults
	)

	for _, hCharacts := range hops {
		if flavorsDes := s.buildCharacteristicsDescription(hCharacts, "sabor"); len(flavorsDes) > 0 {
			flavorDescription += "El lúpulo " + hCharacts.NameID + " aporta " + flavorsDes
		}
		if smellsDes := s.buildCharacteristicsDescription(hCharacts, "aroma"); len(smellsDes) > 0 {
			smellDescription += "El lúpulo " + hCharacts.NameID + " aporta " + smellsDes
		}
		if afterTasteDes := s.buildCharacteristicsDescription(hCharacts, "amargor"); len(afterTasteDes) > 0 {
			afterTasteDescription += "El lúpulo " + hCharacts.NameID + " aporta " + afterTasteDes
		}
	}

	result.SetAfterTasteCharacteristic(afterTasteDescription)
	result.SetFlavorCharacteristic(flavorDescription)
	result.SetSmellCharacteristic(smellDescription)
	return &result
}

// getHops construye una lista de lupulos
// a partir de los datos provistos por el usuario
// y los datos almacenados en base de datos para dichos lupulos
func (s *service) getHops(hops []Hop) ([]HopParams, error) {
	var (
		totalQuantity float64
		hopParams     []HopParams
	)

	for _, h := range hops {
		totalQuantity += h.Quantity
	}

	for _, h := range hops {
		hop, err := s.repo.GetHopByName(h.NameID)
		if err != nil || hop == nil {
			return nil, err
		}

		// ¿qué proporción ocupa del total?
		hopProportion := (totalQuantity * 100) / h.Quantity
		// ¿cuánto contribuye en cd characts segun el tiempo de trabajo?
		charactsWithRealContribution := s.getCharactsWithRealContribution(
			hop.Characteristics(),
			hopProportion,
			h.TimeOfWork)
		hopParam := HopParams{
			NameID:          h.NameID,
			Characteristics: charactsWithRealContribution,
			AlphaAcids:      hop.AlphaAcids(),
			Proportion:      hopProportion,
			Quantity:        h.Quantity,
			TimeOfWork:      h.TimeOfWork,
		}
		hopParams = append(hopParams, hopParam)
	}
	return hopParams, nil
}

// getCharactsWithRealContribution
func (s *service) getCharactsWithRealContribution(characts []ingredients.Characteristic, proportion float64, timeOfWork uint64) []ingredients.Characteristic {
	var characteristics []ingredients.Characteristic
	charactType := s.loadContributionTypeByTimeOfWork(timeOfWork)
	for _, charact := range characts {
		if strings.EqualFold(charact.CharacteristicType, charactType) {
			charact.Contribution = (proportion * charact.Contribution) / 100
			characteristics = append(characteristics, charact)
		}
	}
	return characteristics
}

// loadContributionTypeByTimeOfWork devuelve el tipo de caracteristica
// que mas resalta segun el tiempo de hervor del lupulo
func (s *service) loadContributionTypeByTimeOfWork(timeOfWork uint64) string {
	var (
		arrayTimesOfWork        = []uint64{7, 25, 90}
		arrayCharacteristicType = []string{"aroma", "sabor", "amargor"}
	)

	for i, tm := range arrayTimesOfWork {
		if tm >= timeOfWork {
			return arrayCharacteristicType[i]
		}
	}

	return ""
}

// calculateIBU implementa la formula de calculo del IBU de Glenn Tinseth
func (s *service) calculateIBU(params *Params, addition HopParams) float64 {
	firstFactor := s.greatnessFactor(params.InitialDensity)
	secondFactor := s.boilingTimeFactor(addition.TimeOfWork)
	thirdFactor := s.proportionOfAlphaAcidUsed(addition.AlphaAcids, addition.Quantity)
	divisor := params.WortAmount * 4.15

	if divisor > 0 {
		return (firstFactor * secondFactor * thirdFactor) / divisor
	}
	return 0
}

func (s *service) greatnessFactor(initialDensity float64) float64 {
	if initialDensity <= 0 {
		return 0
	}

	firstFactor := 1.65
	base := 0.000125
	exponent := initialDensity - 1.0
	secondFactor := math.Pow(base, exponent)
	return firstFactor * secondFactor
}

func (s *service) boilingTimeFactor(timeOfWork uint64) float64 {
	exponent := -0.04 * float64(timeOfWork)
	eRaisedExp := math.Exp(exponent)
	dividend := 1 - eRaisedExp

	return dividend
}

func (s *service) proportionOfAlphaAcidUsed(alphaAcids float64, quantity float64) float64 {
	return (alphaAcids / 100.0) * quantity * 1000
}

func (s *service) buildCharacteristicsDescription(hopParam HopParams, characteristicType string) string {
	var (
		characteristicsDescription string
		characteristics            = s.getCharacteristicsByType(hopParam.Characteristics, characteristicType)
		addConnector               = len(characteristics) - 1
	)

	for index, characterisc := range characteristics {
		characteristicsDescription += characterisc.GetDescriptionByProportion(hopParam.Proportion)
		if index < addConnector {
			characteristicsDescription += s.getConnector(index, addConnector)
		}

	}
	return characteristicsDescription
}

func (s *service) getCharacteristicsByType(ccharacts []ingredients.Characteristic, characteristicType string) []ingredients.Characteristic {
	var resultCharacts []ingredients.Characteristic
	for _, characterisc := range ccharacts {
		if strings.EqualFold(characterisc.CharacteristicType, characteristicType) {
			resultCharacts = append(resultCharacts, characterisc)
		}
	}
	return resultCharacts
}

func (s *service) getConnector(index, elements int) string {
	if index+1 == elements {
		return " y "
	}
	return ", "
}

func (s *service) buildResult(ibu float64, flavorResult *FlavorResults) *Results {
	results := &Results{}
	results.SetIBU(ibu)
	results.SetSmellCharacteristic(flavorResult.SmellCCharacteristic())
	results.SetFlavorCharacteristic(flavorResult.FlavorCharacteristic())
	results.SetAfterTasteCharacteristic(flavorResult.AfterTasteCharacteristic())
	return results
}
