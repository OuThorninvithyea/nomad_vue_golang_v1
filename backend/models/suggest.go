package models

type SuggestCard struct {
	Label        string `json:"label"`
	Image        string `json:"image"`
	ImageAlt     string `json:"imageAlt"`
	Rank         int    `json:"rank"`
	Speed        int    `json:"speed"`
	Country      string `json:"country"`
	City         string `json:"city"`
	WeatherEmoji string `json:"weatherEmoji"`
	Temp         int    `json:"temp"`
	TempEmoji    string `json:"tempEmoji"`
	Aqi          int    `json:"aqi"`
	AqiEmoji     string `json:"aqiEmoji"`
	Cost         string `json:"cost"`
}
