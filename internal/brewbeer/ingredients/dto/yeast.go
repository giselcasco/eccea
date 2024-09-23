package dto

import "eccea/internal/brewbeer/ingredients"

type YeastResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	TemperatureMin  float64 `json:"temp_min"`
	TemperatureMax  float64 `json:"temp_max"`
	TimeMin         float64 `json:"time_min"`
	TimeRecommended float64 `json:"time_recommended"`
	TimeMax         float64 `json:"time_max"`
}

func (y *YeastResponse) ToDomain() *ingredients.Yeast {
	return ingredients.NewYeastBuilder().
		Name(y.Name).
		TemperatureRange(y.TemperatureMin, y.TemperatureMax).
		TimeOfWork(y.TimeMin, y.TimeRecommended, y.TimeMax).
		Build()
}
