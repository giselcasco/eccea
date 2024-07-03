package repository

import (
	"bytes"
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/repository/responses"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type ingredientRepo struct {
}

func NewIngredientsRepository() ingredients.Repository {
	return &ingredientRepo{}
}

func (i *ingredientRepo) GetMalt(idMalt string) (*ingredients.Malt, error) {
	return nil, nil
}

func (i *ingredientRepo) GetHop(hopID string) (*ingredients.Hop, error) {
	hop := &responses.Hop{}
	pathHopResources := fmt.Sprintf("internal/resources/hops/%s.json", hopID)
	if i.getResource(pathHopResources, hop); hop.ID == hopID {
		return hop.ToDomain(), nil
	}

	return nil, errors.New("el lupulo no se encuentra en nuestra base de datos")
}

func (i *ingredientRepo) GetYeast(idYeast string) (*ingredients.Yeast, error) {
	return nil, nil
}

func (i *ingredientRepo) getResource(path string, ingredient interface{}) {
	byteArray, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	json.Compact(buffer, byteArray)

	if errUnmarshal := json.Unmarshal([]byte(buffer.String()), ingredient); err != nil {
		panic(errUnmarshal)
	}
}
