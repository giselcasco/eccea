package color_test

import (
	"eccea/internal/brewbeer/estimators/color"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maceration/macerationmocks"
	"eccea/internal/brewbeer/processes/maturation"
	"eccea/internal/brewbeer/processes/maturation/maturationmocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_EstimateFails_When_ServiceFails(t *testing.T) {
	macerationService := &macerationmocks.ServiceMock{}
	maturationService := &maturationmocks.ServiceMock{}
	estimator := color.NewColorImpl(macerationService, maturationService)
	errService := errors.New("sometime error")

	macerationService.On("EstimateColor", &maceration.Params{}).Return(nil, errService)

	_, err := estimator.Estimate(color.Params{maceration.Params{}, maturation.Params{}})

	assert.Error(t, err)
}
func TestShould_EstimateSuccess_When_ValidParams(t *testing.T) {
	macerationService := &macerationmocks.ServiceMock{}
	maturationService := &maturationmocks.ServiceMock{}
	estimator := color.NewColorImpl(macerationService, maturationService)

	// Maceration
	macerationParams := &maceration.Params{
		WortAmount:    50,
		TotalQuantity: 6000,
		MaltAdditions: []maceration.Malt{
			{
				NameID:   "Caramel",
				Quantity: 6000,
			},
		},
	}
	colorResults := &maceration.ColorResults{}
	colorResults.SetColor(1)
	colorResults.SetColorDescription("Pajoso")
	colorResults.SetColorCharacteristic("La malta Caramel aporta intensas tonalidades doradas")
	macerationService.On("EstimateColor", macerationParams).Return(colorResults, nil)

	// Maturation
	maturationParams := &maturation.Params{
		NumberOfDays: 15,
	}
	intensityColor := "Alta claridad, color definido."
	maturationService.On("EstimateColor", maturationParams).Return(intensityColor)

	response, err := estimator.Estimate(color.Params{*macerationParams, *maturationParams})

	assert.NotNil(t, response)
	assert.Equal(t, intensityColor, response.ColorIntensityDescription)
	assert.Equal(t, colorResults.Color(), response.ColorSRM)
	assert.Equal(t, colorResults.ColorDescription(), response.ColorDescription)
	assert.Equal(t, colorResults.ColorCharacteristic(), response.ColorCharacteristics)
	assert.Nil(t, err)
}
