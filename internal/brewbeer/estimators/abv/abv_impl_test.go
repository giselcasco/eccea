package abv_test

import (
	"eccea/internal/brewbeer/estimators/abv"
	"eccea/internal/brewbeer/processes/fermentation"
	"eccea/internal/brewbeer/processes/fermentation/fermentationmocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShould_EstimateSuccess_When_NilParams(t *testing.T) {
	service := &fermentationmocks.ServiceMock{}
	estimator := abv.NewABVImpl(service)
	abvValue := 5.9
	service.On("CalculateABV", &fermentation.Params{}).Return(abvValue, nil)

	response, err := estimator.EstimateAlcoholByVolume(abv.Params{})

	assert.Equal(t, response, abvValue)
	assert.Equal(t, err, nil)
}

func float(f float64) {
	panic("unimplemented")
}

func TestShould_EstimateFails_When_ServiceFails(t *testing.T) {
	service := &fermentationmocks.ServiceMock{}
	estimator := abv.NewABVImpl(service)
	errorMock := errors.New("something wrong")

	service.On("Do", mock.Anything, mock.Anything).Return(nil, errorMock)

	response, err := estimator.EstimateAlcoholByVolume(abv.Params{})

	assert.NotNil(t, response)
	assert.Error(t, err)
}
