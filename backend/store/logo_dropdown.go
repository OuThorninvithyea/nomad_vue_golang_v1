package store

import "backend/models"

var LogoDropdown = []models.LogoDropdownSection{
	{
		SectionClass: "header-logo-clicked-info",
		Items: []models.LogoDropdownItem{
			{Text: "Join Nomads.com", Emoji: "🌍"},
			{Text: "Log in", Emoji: "👋"},
		},
	},
	{
		Title: "GENERAL",
		Items: []models.LogoDropdownItem{
			{Text: "Log in", Emoji: "👋"},
			{Text: "Frontpage", Emoji: "🏠"},
			{Text: "Dark mode", Emoji: "🌗"},
			{Text: "Your favorites", Emoji: "❤️"},
			{Text: "Nomad insurance", Emoji: "🚑"},
		},
	},
	{
		Title: "COMMUNITY",
		Items: []models.LogoDropdownItem{
			{Text: "Dating app", Emoji: "❤️"},
			{Text: "Friend finder", Emoji: "🤝"},
			{Text: "Chat", Emoji: "💬"},
			{Text: "Members map", Emoji: "🌍"},
			{Text: "Host meetup", Emoji: "💈", Badge: "NEW"},
			{Text: "Attend meetup", Emoji: "🍹"},
		},
	},
	{
		Title: "TOOLS",
		Items: []models.LogoDropdownItem{
			{Text: "Explore", Emoji: "🏕️"},
			{Text: "NomadGPT", Emoji: "🎒"},
			{Text: "Vote on photos", Emoji: "📸"},
			{Text: "Residence calendar", Emoji: "🗓️"},
			{Text: "FIRE calculator", Emoji: "💸"},
			{Text: "Climate finder", Emoji: "🌤️"},
			{Text: "Fastest growing", Emoji: "💥"},
			{Text: "Nomad stats", Emoji: "📊", Badge: "NEW"},
			{Text: "State of remote work", Emoji: "🧪", Badge: "NEW"},
			{Text: "History of nomads", Emoji: "📜", Badge: "NEW"},
			{Text: "Network graph", Emoji: "🕸️"},
			{Text: "Fastest internet", Emoji: "🔌", Tooltip: ">10mbps"},
			{Text: "Best place now", Emoji: "🏆"},
			{Text: "Random place", Emoji: "🔮"},
			{Text: "Random good place", Emoji: "🔮"},
		},
	},
	{
		Title: "HELP",
		Items: []models.LogoDropdownItem{
			{Text: "Ideas + bugs", Emoji: "💡"},
			{Text: "Changelog", Emoji: "🚀"},
			{Text: "Merch", Emoji: "🛍️"},
			{Text: "FAQ & Help", Emoji: "🛟"},
			{Text: "TOS & Privacy policy", Emoji: "👩‍💼"},
		},
	},
	{
		Title: "OTHER PROJECTS",
		Items: []models.LogoDropdownItem{
			{Text: "Hotel List", Emoji: "🏥"},
			{Text: "Remote jobs", Emoji: "🛰️"},
			{Text: "Airline List", Emoji: "✈️"},
			{Text: "Neighborhoods", Emoji: "🗺️"},
			{Text: "Pieter", Emoji: "💾"},
			{Text: "fly.pieter.com", Emoji: "✈️"},
			{Text: "Coworkations", Emoji: "🏝️"},
			{Text: "Nomad Cafe", Emoji: "☕"},
			{Text: "Interior AI", Emoji: "🏡"},
			{Text: "Photo AI", Emoji: "📸"},
			{Text: "eu/acc", Emoji: "🇪🇺"},
			{Text: "JSON.pub", Emoji: "📄"},
		},
	},
}
