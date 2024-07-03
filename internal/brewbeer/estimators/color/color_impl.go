package color

import (
	"eccea/internal/brewbeer/characteristic"
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
func (c *colorImpl) Estimate(params Params) *Estimation {
	results := c.macerationService.EstimateColor(&params.Maceration)
	return NewEstimation(results.Color(), []characteristic.Color{})
}
