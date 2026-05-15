package store

import "backend/models"

var Suggest = models.SuggestCard{
	Label:        "nomad.com suggestion",
	Image:        "/images/phnopenh.webp",
	ImageAlt:     "Porto",
	Rank:         1,
	Speed:        65,
	Country:      "Portugal",
	City:         "Porto",
	WeatherEmoji: "🌤️",
	Temp:         19,
	TempEmoji:    "😊",
	Aqi:          22,
	AqiEmoji:     "😷",
	Cost:         "$2200/mo",
}
