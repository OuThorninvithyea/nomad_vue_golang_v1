package models

type City struct {
	Image        string   `json:"image"`
	Alt          string   `json:"alt"`
	Rank         int      `json:"rank"`
	Speed        int      `json:"speed"`
	Label        string   `json:"label"`
	Country      string   `json:"country"`
	City         string   `json:"city"`
	WeatherEmoji string   `json:"weatherEmoji"`
	Temp         int      `json:"temp"`
	TempEmoji    string   `json:"tempEmoji"`
	Aqi          int      `json:"aqi"`
	AqiEmoji     string   `json:"aqiEmoji"`
	Cost         string   `json:"cost"`
	Ratings      []Rating `json:"ratings"`
}

type Rating struct {
	Emoji    string `json:"emoji"`
	Label    string `json:"label"`
	BarClass string `json:"barClass"`
}
