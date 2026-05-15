package store

import "backend/models"

var TrustedCompany = models.TrustedCompany{
	Header: "as seen on",
	Companies: []models.Company{
		{Src: "/images/newyorktimes.png", Alt: "New York Times"},
		{Src: "/images/ft.png", Alt: "Financial Times"},
		{Src: "/images/bbc.png", Alt: "BBC"},
		{Src: "/images/cnn.png", Alt: "CNN"},
		{Src: "/images/usa-today.png_8.png", Alt: "USA Today"},
		{Src: "/images/cnbc.png", Alt: "CNBC"},
		{Src: "/images/guardian.png", Alt: "The Guardian"},
		{Src: "/images/politico.png", Alt: "Politico"},
	},
}
