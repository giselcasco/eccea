package beer

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/fermentation"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
	"strings"
)

type beerImpl struct {
	macerationService   maceration.Service
	boilingService      boiling.Service
	fermentationService fermentation.Service
	maturationService   maturation.Service
}

func NewBeerImpl(macerationService maceration.Service, boilingService boiling.Service, fermentationService fermentation.Service, maturationService maturation.Service) Beer {
	return &beerImpl{
		macerationService:   macerationService,
		boilingService:      boilingService,
		fermentationService: fermentationService,
		maturationService:   maturationService,
	}
}
func (b *beerImpl) Estimate(params Params) (*Estimation, error) {
	macerationResults, macerationErr := b.macerationService.EstimateBeer(params.Maceration)
	if macerationErr != nil {
		return nil, macerationErr
	}

	boilingResults, boilingErr := b.boilingService.EstimateBeer(params.Boiling)
	if boilingErr != nil {
		return nil, boilingErr
	}

	fermentationResults, fermentationErr := b.fermentationService.EstimateBeer(params.Fermentation)
	if fermentationErr != nil {
		return nil, fermentationErr
	}

	maturationResults := b.maturationService.EstimateBeer(params.Maturation)
	return b.buildResults(boilingResults, fermentationResults, macerationResults, maturationResults), nil
}

func (b *beerImpl) buildResults(
	boilingResults *boiling.Results,
	fermentationResult *fermentation.Results,
	macerationResults *maceration.Results,
	maturationResults maturation.Results) *Estimation {
	flavorCharacteristics := []string{
		boilingResults.FlavorCharacteristic(),
		fermentationResult.FlavorCharacteristic(),
		maturationResults.FlavorMaturation(),
	}
	smellCharacteristics := []string{
		boilingResults.SmellCCharacteristic(),
		fermentationResult.SmellCCharacteristic(),
	}
	afterTasteCharacteristics := []string{
		boilingResults.AfterTasteCharacteristic(),
	}
	colorCharacteristics := []string{
		fermentationResult.ColorCharacteristic(),
		maturationResults.ColorMaturation(),
	}
	return &Estimation{
		IBU:                       boilingResults.IBU(),
		ABV:                       fermentationResult.ABV(),
		ColorDescription:          macerationResults.ColorDescription(),
		ColorSRM:                  macerationResults.Color(),
		ColorCharacteristics:      strings.Join(colorCharacteristics, ", "),
		FlavorCharacteristics:     strings.Join(flavorCharacteristics, ", "),
		SmellCharacteristics:      strings.Join(smellCharacteristics, ", "),
		AfterTasteCharacteristics: strings.Join(afterTasteCharacteristics, ", "),
	}
}
