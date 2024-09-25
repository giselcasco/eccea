package dto

type CharacteristicResponse struct {
	ID           int64   `json:"id"`
	Description  string  `json:"description"`
	AdjetiveID   int     `json:"adjetive_id"`
	Type         string  `json:"type"`
	Contribution float64 `json:"contribution"`
}
