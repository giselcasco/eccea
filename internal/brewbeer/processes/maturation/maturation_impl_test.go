package maturation_test

import (
	"eccea/internal/brewbeer/processes/maturation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	var dataSet = []struct {
		nameTest string
		params   *maturation.Params
		result   string
	}{
		{
			nameTest: "maturation success when nil params provided",
			params:   &maturation.Params{},
			result:   "",
		},
		{
			nameTest: "maturation success when valid params provided",
			params: &maturation.Params{
				NumberOfDays: 30,
			},
			result: "Alta claridad, color definido.",
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			maturationProcess := maturation.NewService()

			response := maturationProcess.EstimateColor(data.params)

			assert.Equal(t, data.result, response)
		})
	}
}
