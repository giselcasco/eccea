package flavor

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
	"strings"
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
	macerationResults, macerationErr := c.macerationService.EstimateFlavor(&params.Maceration)
	if macerationErr != nil {
		return nil, macerationErr
	}
	boilingResults, boilingErr := c.boilingService.EstimateFlavor(&params.Boiling)
	if boilingErr != nil {
		return nil, boilingErr
	}

	return NewEstimation(
		c.joinResult(macerationResults.FlavorCharacteristic(), boilingResults.FlavorCharacteristic()),
		c.joinResult(macerationResults.SmellCharacteristic(), boilingResults.SmellCCharacteristic()),
		boilingResults.AfterTasteCharacteristic()), nil
}

func (c *flavorImpl) joinResult(characteristics string, moreCharacteristics string) string {
	return strings.Join([]string{characteristics, moreCharacteristics}, ", ")
}
