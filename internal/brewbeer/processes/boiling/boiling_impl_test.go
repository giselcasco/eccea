package boiling_test

import (
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/brewbeer/ingredients/repositorymocks"
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/repository/responses"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShould_DoSuccess_When_ParamsOK(t *testing.T) {
	hop := &responses.Hop{AlphaAcids: 5.4}

	var dataSet = []struct {
		nameTest     string
		params       *boiling.Params
		useCaseKey   string
		result       *boiling.Results
		hopResult    *ingredients.Hop
		repoError    error
		boilingError error
		resultError  error
	}{
		{
			nameTest: "Boiling success when nil params provided",
			params: &boiling.Params{
				HopAdditions: []boiling.HopAdditions{
					{
						ID: "2",
					},
				},
			},
			useCaseKey: "ibu",
			hopResult:  hop.ToDomain("2"),
			result:     &boiling.Results{},
		},
		// TODO test with valid values
		{
			nameTest: "Boiling fail when get hop fail",
			params: &boiling.Params{
				HopAdditions: []boiling.HopAdditions{
					{
						ID: "2",
					},
				},
			},
			useCaseKey:  "ibu",
			repoError:   errors.New("something error"),
			resultError: errors.New("something error"),
		},
	}

	for _, data := range dataSet {
		t.Run(data.nameTest, func(t *testing.T) {
			repo := &repositorymocks.RepositoryMock{}
			boiling := boiling.NewService(repo)

			repo.
				On("GetHop", mock.Anything).
				Return(data.hopResult, data.repoError)

			response, err := boiling.Do(data.params, data.useCaseKey)

			assert.Equal(t, data.result, response)
			assert.Equal(t, data.resultError, err)
		})
	}
}
