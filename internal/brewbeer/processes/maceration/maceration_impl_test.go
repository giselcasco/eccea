package maceration_test

import (
	"eccea/internal/brewbeer/ingredients/repositorymocks"
	"eccea/internal/brewbeer/processes/maceration"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	var dataSet = []struct {
		nameTest string
		params   *maceration.Params
		result   *maceration.ColorResults
		err      error
	}{
		{
			nameTest: "Fermentation success when nil params provided",
			params:   nil,
			result:   nil,
			err:      errors.New("nil params error"),
		},
		{
			nameTest: "Fermentation success when valid params provided",
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
			result: &maceration.ColorResults{},
			err:    nil,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			repo := &repositorymocks.RepositoryMock{}
			macerationProcess := maceration.NewService(repo)

			response, err := macerationProcess.EstimateColor(data.params)

			assert.Equal(t, data.result, response)
			assert.Equal(t, data.err, err)
		})
	}
}
