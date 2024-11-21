package models

type Settings struct {
	City           string   `json:"city"`
	ExcludedPlaces []string `json:"excluded_places"`
	ExcludedDays   []string `json:"excluded_days"`
}
