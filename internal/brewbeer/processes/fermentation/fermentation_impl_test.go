package fermentation_test

import (
	"eccea/internal/brewbeer/processes/fermentation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	var dataSet = []struct {
		nameTest       string
		initialDensity float64
		finalDensity   float64
		expectedABV    float64
		expectedError  bool
	}{
		{
			nameTest:       "Fermentation success when nil params provided",
			initialDensity: 0,
			finalDensity:   0,
			expectedABV:    0,
			expectedError:  false,
		},
		{
			nameTest:       "Fermentation success when valid params provided",
			initialDensity: 1.050,
			finalDensity:   1.010,
			expectedABV:    5.25,
			expectedError:  false,
		},
		{
			nameTest:       "Fermentation succes when initial density is nil",
			initialDensity: 0,
			finalDensity:   1.010,
			expectedABV:    0,
			expectedError:  true,
		},
		{
			nameTest:       "Fermentation fail when initial density is negative",
			initialDensity: -1.050,
			finalDensity:   1.010,
			expectedABV:    0,
			expectedError:  true,
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			fermentationProcess := fermentation.NewService()

			response, err := fermentationProcess.CalculateAlcoholByVolume(data.initialDensity, data.finalDensity)

			assert.Equal(t, data.expectedError, err == nil)
			assert.Equal(t, data.expectedABV, response)
		})
	}
}
