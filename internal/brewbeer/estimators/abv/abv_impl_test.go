package ibu_test

import (
	"eccea/internal/brewbeer/estimators/ibu"
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/boiling/boilingmocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShould_EstimateSuccess_When_NilParams(t *testing.T) {
	service := &fermentationmocks.ServiceMock{}
	estimator := abv.NewABVImpl(service)

	service.On("Do", &fermentation.Params{}, "abv").Return(&fermentation.Results{}, nil)

	response, err := estimator.Estimate(abv.Params{})

	assert.Equal(t, response, abv.Estimation{})
	assert.Equal(t, err, nil)
}

func TestShould_EstimateFails_When_ServiceFails(t *testing.T) {
	service := &fermentationmocks.ServiceMock{}
   	estimator := abv.NewABVImpl(service)
    errorMock := errors.New("something wrong")

	service.On("Do", mock.Anything, mock.Anything).Return(nil, errorMock)

	response, err := estimator.Estimate(abv.Params{})

	assert.NotNil(t, response)
	assert.Error(t, err)
}
