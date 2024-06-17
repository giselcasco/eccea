package fermentation_test

import (
	"eccea/internal/brewbeer/processes/fermentation"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	useCaseKey = "abv"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	var dataSet = []struct {
		nameTest string
		params   *fermentation.Params
		abv      float32
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
			abv: 5.9062443,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			fermentationProcess := fermentation.NewService()

			response := fermentationProcess.Do(data.params, useCaseKey)

			assert.Equal(t, data.abv, response.ABV())
		})
	}
}
