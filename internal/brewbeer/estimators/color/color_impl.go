package color

import (
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

// colorImpl es la implementacion para el caso de uso de calculo de IBU
type colorImpl struct {
	macerationService maceration.Service
	maturationService maturation.Service
}

func NewColorImpl(macerationServ maceration.Service, maturationServ maturation.Service) Color {
	return &colorImpl{
		macerationService: macerationServ,
		maturationService: maturationServ,
	}
}

// Estimate es la implementación para el calculo de la estimación del color final
// en unidades de SRM y caracteristicas asociadas al color
func (c *colorImpl) Estimate(params Params) (*Estimation, error) {
	results, err := c.macerationService.EstimateColor(&params.Maceration)
	if err != nil {
		return nil, err
	}
	colorIntensityDescription := c.maturationService.EstimateColor(&params.Maturation)
	return NewEstimation(
		results.Color(),
		results.ColorDescription(),
		colorIntensityDescription,
		results.ColorCharacteristic()), nil
}
