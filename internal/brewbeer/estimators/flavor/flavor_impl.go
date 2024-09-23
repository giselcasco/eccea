package flavor

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

// flavorImpl es la implementacion para el caso de uso de estimación de caracteristicas del sabor
type flavorImpl struct {
	macerationService maceration.Service
	boilingService    boiling.Service
	maturationService maturation.Service
}

func NewFlavorImpl(macerationServ maceration.Service,
	boilingServ boiling.Service,
	maturationServ maturation.Service) Flavor {
	return &flavorImpl{
		macerationService: macerationServ,
		boilingService:    boilingServ,
		maturationService: maturationServ,
	}
}

// Estimate es la implementación para la estimación de caracteristicas del sabor
func (c *flavorImpl) Estimate(params Params) (*Estimation, error) {
	macerationResults, macarationErr := c.macerationService.EstimateFlavor(&params.Maceration)
	if macarationErr != nil {
		return nil, macarationErr
	}
	boilingResults, boilingErr := c.boilingService.EstimateIBU(&params.Boiling)
	if boilingErr != nil {
		return nil, boilingErr
	}
	// TODO get characterists AfterTasteChar
	return NewEstimation(
		macerationResults.FlavorCharacteristic(),
		macerationResults.SmellCharacteristic(),
		macerationResults.SmellCharacteristic()), nil
}
