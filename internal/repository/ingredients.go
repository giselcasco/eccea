package repository

import (
	"context"
	"eccea/internal/brewbeer/characteristic"
	"eccea/internal/brewbeer/ingredients"
	"fmt"
)

type (
	HopResponse struct {
		Name       string  `json:"name"`
		AlphaAcids float64 `json:"alpha_acids"`
		BetaAcids  float64 `json:"beta_acids"`
	}

	MaltResponse struct {
		Name                   string  `json:"name"`
		ColorSRM               uint64  `json:"color"`
		TemperatureMin         uint64  `json:"temp_min"`
		TemperatureRecommended uint64  `json:"temp_recommended"`
		TemperatureMax         uint64  `json:"temp_max"`
		ExtractFineGrind       float64 `json:"extract_fine_grind"`
		ExtractCoarseGrind     float64 `json:"extract_coarse_grind"`
		DiastaticPower         float64 `json:"diastatic_power"`
	}

	YeastResponse struct {
		Name                   string  `json:"name"`
		TemperatureMin         uint64  `json:"temp_min"`
		TemperatureRecommended uint64  `json:"temp_recommended"`
		TemperatureMax         uint64  `json:"temp_max"`
		TimeMin                float64 `json:"time_min"`
		TimeMax                float64 `json:"time_max"`
	}

	CharacteristicResponse struct {
		ID          int64  `json:"id"`
		AdverbType  uint64 `json:"adverb_type"`
		Description string `json:"description"`
	}

	PropertiesResponse struct {
		CharacteristicID int64  `json:"characteristic_id"`
		IngredientName   string `json:"ingredient_name"`
		Type             string `json:"type"`
	}
)

func (h HopResponse) String() string {
	return fmt.Sprintf(" -> Name: %s , -> AlphaAcids: %f, -> BetaAcids: %f", h.Name, h.AlphaAcids, h.BetaAcids)
}

func (m MaltResponse) String() string {
	return fmt.Sprintf(" -> Name: %s , -> ColorSRM: %d", m.Name, m.ColorSRM)
}

func (m MaltResponse) ToDomain() *ingredients.Malt {
	maltBuilder := ingredients.NewMaltBuilder().ColorSRM(float64(m.ColorSRM))
	return fmt.Sprintf(" -> Name: %s , -> ColorSRM: %d", m.Name, m.ColorSRM)
}

func (y YeastResponse) String() string {
	return fmt.Sprintf(" -> Name: %s", y.Name)
}

type Reader interface {
	GetMalt(ctx context.Context, maltName string) (*ingredients.Malt, error)
	GetHop(ctx context.Context, hopName string) (*ingredients.Hop, error)
	GetYeast(ctx context.Context, yeastName string) (*ingredients.Yeast, error)

	GetColorCharacteristics(ctx context.Context, ingredientName string) (*characteristic.Color, error)
	GetFlavorCharacteristics(ctx context.Context, ingredientName string) (*characteristic.Flavor, error)
}
