package repository

import (
	"context"
	"eccea/internal/brewbeer/ingredients"
	"slices"
)

type (
	HopResponse struct {
		ID              string                   `json:"id"`
		Name            string                   `json:"name"`
		AlphaAcids      float64                  `json:"alpha_acids"`
		BetaAcids       float64                  `json:"beta_acids"`
		Characteristics []CharacteristicResponse `json:"characteristics"`
	}

	MaltResponse struct {
		ID                     string                   `json:"id"`
		Name                   string                   `json:"name"`
		ColorSRM               float64                  `json:"color"`
		TemperatureMin         uint64                   `json:"temp_min"`
		TemperatureRecommended uint64                   `json:"temp_recommended"`
		TemperatureMax         uint64                   `json:"temp_max"`
		ExtractFineGrind       float64                  `json:"extract_fine_grind"`
		ExtractCoarseGrind     float64                  `json:"extract_coarse_grind"`
		DiastaticPower         float64                  `json:"diastatic_power"`
		Characteristics        []CharacteristicResponse `json:"characteristics"`
	}

	YeastResponse struct {
		ID              string  `json:"id"`
		Name            string  `json:"name"`
		TemperatureMin  float64 `json:"temp_min"`
		TemperatureMax  float64 `json:"temp_max"`
		TimeMin         float64 `json:"time_min"`
		TimeRecommended float64 `json:"time_recommended"`
		TimeMax         float64 `json:"time_max"`
	}

	CharacteristicResponse struct {
		ID          int64    `json:"id"`
		Description string   `json:"description"`
		AdjetiveID  uint64   `json:"adjetive_id"`
		Types       []string `json:"types"`
	}
)

func (y *YeastResponse) ToDomain() *ingredients.Yeast {
	return ingredients.NewYeastBuilder().
		Name(y.Name).
		TemperatureRange(y.TemperatureMin, y.TemperatureMax).
		TimeOfWork(y.TimeMin, y.TimeRecommended, y.TimeMax).
		Build()
}

func (h *HopResponse) ToDomain() *ingredients.Hop {
	scharacts := []ingredients.Smell{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "aroma") {
			scharact := ingredients.Smell{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			scharacts = append(scharacts, scharact)
		}
	}

	fcharacts := []ingredients.Flavor{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "sabor") {
			fcharact := ingredients.Flavor{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			fcharacts = append(fcharacts, fcharact)
		}
	}

	acharacts := []ingredients.AfterTaste{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "retrogusto") {
			acharact := ingredients.AfterTaste{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			acharacts = append(acharacts, acharact)
		}
	}

	return ingredients.NewHopBuilder().
		Name(h.Name).
		FlavorCharacteristics(fcharacts).
		SmellCharacteristics(scharacts).
		AfterTasteCharacteristics(acharacts).
		AlphaAcids(h.AlphaAcids).
		BetaAcids(h.BetaAcids).
		Build()
}

func (m *MaltResponse) ToDomain() *ingredients.Malt {
	ccharacts := []ingredients.Color{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "color") {
			ccharact := ingredients.Color{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			ccharacts = append(ccharacts, ccharact)
		}
	}

	scharacts := []ingredients.Smell{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "aroma") {
			scharact := ingredients.Smell{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			scharacts = append(scharacts, scharact)
		}
	}

	fcharacts := []ingredients.Flavor{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "sabor") {
			fcharact := ingredients.Flavor{
				Type:        ingredients.TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			fcharacts = append(fcharacts, fcharact)
		}
	}

	return ingredients.NewMaltBuilder().
		ColorSRM(m.ColorSRM).
		ColorCharacteristics(ccharacts).
		SmellCharacteristics(scharacts).
		FlavorCharacteristics(fcharacts).
		Name(m.Name).
		Build()
}

type Reader interface {
	GetMalt(ctx context.Context, maltName string) (*ingredients.Malt, error)
	GetHop(ctx context.Context, hopName string) (*ingredients.Hop, error)
	GetYeast(ctx context.Context, yeastName string) (*ingredients.Yeast, error)
}
