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

	service.On("Do", &boiling.Params{}, "ibu").Return(&boiling.Results{}, nil)

	response, err := estimator.Estimate(ibu.Params{})

	assert.Equal(t, response, ibu.Estimation{})
	assert.Equal(t, err, nil)
}
func TestShould_EstimateFails_When_ServiceFails(t *testing.T) {
	service := &boilingmocks.ServiceMock{}
	estimator := ibu.NewIBUImpl(service)
	errorMock := errors.New("something wrong")

	service.On("Do", mock.Anything, mock.Anything).Return(nil, errorMock)

	response, err := estimator.Estimate(ibu.Params{})

	assert.NotNil(t, response)
	assert.Error(t, err)
}
