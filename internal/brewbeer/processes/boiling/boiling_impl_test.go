package boiling_test

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/repository/repositorymocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {

	var dataSet = []struct {
		nameTest     string
		params       *boiling.Params
		ibu          float64
		boilingError error
		resultError  error
	}{
		{
			nameTest: "Boiling success when nil params provided",
			params:   &boiling.Params{},
		},
		{
			nameTest: "Boiling success when valid params provided",
			params: &boiling.Params{
				InitialDensity: 1.024,
				WortAmount:     50,
				HopAdditions: []boiling.Hop{
					{
						AlphaAcids: 14.0,
						TimeOfWork: 30,
						Quantity:   200,
					},
				},
			},
			ibu: 125.40250396754854,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			repo := repositorymocks.NewRepositoryMock(t)
			boiling := boiling.NewService(repo)

			response, err := boiling.EstimateIBU(data.params)

			assert.Equal(t, data.ibu, response)
			assert.Equal(t, data.resultError, err)
		})
	}
}
