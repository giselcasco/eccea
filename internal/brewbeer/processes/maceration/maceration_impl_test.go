package maceration_test

import (
	"eccea/internal/brewbeer/characteristic"
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/brewbeer/ingredients/repositorymocks"
	"eccea/internal/brewbeer/processes/maceration"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	maltBuilder := ingredients.NewMaltBuilder()
	maltBuilder.Name("Caramel")
	maltBuilder.ColorSRM(10)
	maltBuilder.ColorCharacteristics([]characteristic.Color{{
		Type:        1,
		Description: characteristic.GoldColor,
	}})

	colorResults := &maceration.ColorResults{}
	colorResults.SetColor(1)
	colorResults.SetColorDescription("Pajoso")
	colorResults.SetColorCharacteristic("La malta Caramel aporta intensas tonalidades doradas")

	var dataSet = []struct {
		nameTest      string
		params        *maceration.Params
		resultService *maceration.ColorResults
		resultRepo    *ingredients.MaltBuilder
		errService    error
		errRepo       error
	}{
		{
			nameTest:      "maceration fails when nil params provided",
			params:        nil,
			resultService: nil,
			errService:    errors.New("nil params error"),
		},
		{
			nameTest: "maceration fails when get malt fails",
			params: &maceration.Params{
				WortAmount:    50,
				TotalQuantity: 6000,
				MaltAdditions: []maceration.Malt{
					{
						NameID:   "Caramel",
						Quantity: 6000,
					},
				},
			},
			resultRepo:    maltBuilder,
			errRepo:       errors.New("error when get malt"),
			resultService: nil,
			errService:    errors.New("error when get malt"),
		},
		{
			nameTest: "maceration success when valid params provided",
			params: &maceration.Params{
				WortAmount:    50,
				TotalQuantity: 6000,
				MaltAdditions: []maceration.Malt{
					{
						NameID:   "Caramel",
						Quantity: 6000,
					},
				},
			},
			resultRepo:    maltBuilder,
			resultService: colorResults,
			errService:    nil,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			repo := &repositorymocks.RepositoryMock{}
			macerationProcess := maceration.NewService(repo)
			if data.params != nil && len(data.params.MaltAdditions) > 0 {
				repo.
					On("GetMalt", data.params.MaltAdditions[0].NameID).
					Return(data.resultRepo.Build(), data.errRepo)
			}

			response, err := macerationProcess.EstimateColor(data.params)

			assert.Equal(t, data.resultService, response)
			assert.Equal(t, data.errService, err)
		})
	}
}
