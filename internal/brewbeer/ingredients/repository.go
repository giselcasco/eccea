package ingredients

import (
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
		ID                 string  `json:"id"`
		Name               string  `json:"name"`
		ColorSRM           float64 `json:"color"`
		TemperatureMin     float64 `json:"temp_min"`
		TemperatureMax     float64 `json:"temp_max"`
		ExtractFineGrind   float64 `json:"extract_fine_grind"`
		ExtractCoarseGrind float64 `json:"extract_coarse_grind"`
		// DiastaticPower bajo o nulo en maltas caramelizadas o tostadas
		DiastaticPower  float64                  `json:"diastatic_power"`
		Characteristics []CharacteristicResponse `json:"characteristics"`
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

//go:generate mockery --name=Repository --structname=RepositoryMock --case underscore --output repositorymocks  --outpkg repositorymocks
type Reader interface {
	GetMaltByName(maltName string) (*Malt, error)
	GetHopByName(hopName string) (*Hop, error)
	GetYeastByName(yeastName string) (*Yeast, error)
}

func (y *YeastResponse) ToDomain() *Yeast {
	return NewYeastBuilder().
		Name(y.Name).
		TemperatureRange(y.TemperatureMin, y.TemperatureMax).
		TimeOfWork(y.TimeMin, y.TimeRecommended, y.TimeMax).
		Build()
}

func (h *HopResponse) ToDomain() *Hop {
	scharacts := []Smell{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "aroma") {
			scharact := Smell{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			scharacts = append(scharacts, scharact)
		}
	}

	fcharacts := []Flavor{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "sabor") {
			fcharact := Flavor{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			fcharacts = append(fcharacts, fcharact)
		}
	}

	acharacts := []AfterTaste{}
	for _, cc := range h.Characteristics {
		if slices.Contains(cc.Types, "retrogusto") {
			acharact := AfterTaste{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			acharacts = append(acharacts, acharact)
		}
	}

	return NewHopBuilder().
		Name(h.Name).
		FlavorCharacteristics(fcharacts).
		SmellCharacteristics(scharacts).
		AfterTasteCharacteristics(acharacts).
		AlphaAcids(h.AlphaAcids).
		BetaAcids(h.BetaAcids).
		Build()
}

func (m *MaltResponse) ToDomain() *Malt {
	ccharacts := []Color{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "color") {
			ccharact := Color{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			ccharacts = append(ccharacts, ccharact)
		}
	}

	scharacts := []Smell{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "aroma") {
			scharact := Smell{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			scharacts = append(scharacts, scharact)
		}
	}

	fcharacts := []Flavor{}
	for _, cc := range m.Characteristics {
		if slices.Contains(cc.Types, "sabor") {
			fcharact := Flavor{
				Type:        TypeAdjetives(cc.AdjetiveID),
				Description: cc.Description,
			}
			fcharacts = append(fcharacts, fcharact)
		}
	}

	return NewMaltBuilder().
		ColorSRM(m.ColorSRM).
		ColorCharacteristics(ccharacts).
		SmellCharacteristics(scharacts).
		FlavorCharacteristics(fcharacts).
		TemperatureRange(m.TemperatureMin, m.TemperatureMax).
		Name(m.Name).
		Build()
}
