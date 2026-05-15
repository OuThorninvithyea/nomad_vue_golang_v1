package store

import "backend/models"

var HoverCards = []models.HoverCard{
	{
		Quote:          "Nomads.com ranks destinations that are accommodating to digital nomads, based on factors like cost of living, internet speed and weather",
		Image:          "/images/newyorktimes.png",
		ImageAlt:       "new-york-times",
		AnimationClass: "turn-right",
	},
	{
		Quote:          "The rankings of Nomads.com's cities are constantly in flux (all the data is refreshed in real-time based on user input)",
		Image:          "/images/bbc.png",
		ImageAlt:       "bbc",
		AnimationClass: "zoomin-card",
	},
	{
		Quote:          "Nomads.com is a Kayak-like aggregator for potential work destinations, ranking internet, price, and safety",
		Image:          "/images/techcrunch.png",
		ImageAlt:       "techcrunch",
		AnimationClass: "turn-left",
	},
}
