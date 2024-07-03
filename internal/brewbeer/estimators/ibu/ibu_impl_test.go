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
	service := &boilingmocks.ServiceMock{}
	estimator := ibu.NewIBUImpl(service)
	ibuValue := 0.1

	service.On("EstimateIBU", &boiling.Params{}).Return(ibuValue, nil)

	response, err := estimator.Estimate(ibu.Params{})

	assert.Equal(t, response, ibuValue)
	assert.Equal(t, err, nil)
}
func TestShould_EstimateFails_When_ServiceFails(t *testing.T) {
	service := &boilingmocks.ServiceMock{}
	estimator := ibu.NewIBUImpl(service)
	errorMock := errors.New("something wrong")

	service.On("EstimateIBU", mock.Anything, mock.Anything).Return(0, errorMock)

	response, err := estimator.Estimate(ibu.Params{})

	assert.NotNil(t, response)
	assert.Error(t, err)
}
