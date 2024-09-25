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
	for _, hopAddition := range params.HopAdditions {
		ibu += s.calculateIBU(params, hopAddition)
	}

	return ibu, nil
}

/*
EstimateIBU es el metodo que calcula el ibu a partir de los valores en los parametros
ingresados por el usuario.
*/
func (s *service) EstimateFlavor(params *Params) (*FlavorResults, error) {
	_, herr := s.getHops(params.HopAdditions)
	if herr != nil {
		return nil, herr
	}
	// TODO cargar perfil sensorial y
	// calcular porcentaje de contr¡bucion para aroma, amargor y aroma
	// y segun eso armar definición del aporte de cada lupulo

	results := FlavorResults{}
	results.SetFlavorCharacteristic("ccc")
	return nil, nil
}

// getHop busca el lupudo en la lista de lupulos "hops" cuyo ID corresponda con "idHop"
// opcional para busqueda en db del alphaAcids del hop
func getHop(hops []ingredients.Hop, idHop string) *ingredients.Hop {
	for _, hop := range hops {
		if strings.EqualFold(hop.Name(), idHop) {
			return &hop
		}
	}
	return nil
}

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
		if err != nil {
			return nil, err
		}

		hopProportion := (totalQuantity * 100) / h.Quantity
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
		}
		hopParams = append(hopParams, hopParam)
	}
	return hopParams, nil
}

func (s *service) getCharactsWithRealContribution(characts []ingredients.Characteristic, proportion float64, timeOfWork uint64) []ingredients.Characteristic {
	var characteristics []ingredients.Characteristic
	charactType := s.loadContributionTypeByTimeOfWork(timeOfWork)
	for _, charact := range characts {
		if strings.EqualFold(charact.CharacteristicType, charactType) {
			characteristics = append(characteristics, charact)
		}
	}
	return characteristics
}

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
func (s *service) calculateIBU(params *Params, addition Hop) float64 {
	firstFactor := s.greatnessFactor(params.InitialDensity)
	secondFactor := s.boilingTimeFactor(addition.TimeOfWork)
	thirdFactor := s.proportionOfAlphaAcidUsed(addition.AlphaAcids, addition.Quantity)
	divisor := float64(params.WortAmount) * 4.15

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
