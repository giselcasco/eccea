package fermentation_test

import (
	"eccea/internal/brewbeer/processes/fermentation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	var dataSet = []struct {
		nameTest string
		params   *fermentation.Params
		abv      float64
	}{
		{
			nameTest: "Fermentation success when nil params provided",
			params:   &fermentation.Params{},
			abv:      0,
		},
		{
			nameTest: "Fermentation success when valid params provided",
			params: &fermentation.Params{
				InitialDensity: 1.055,
				FinalDensity:   1.010,
			},
			abv: 5.906249999999991,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			fermentationProcess := fermentation.NewService()

			response := fermentationProcess.CalculateABV(data.params)

			assert.Equal(t, data.abv, response)
		})
	}
}
