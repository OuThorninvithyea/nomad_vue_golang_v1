package store

import "backend/models"

var SearchFilters = []models.SearchFilterItem{
	{Type: "header", Text: "WHAT"},
	{Type: "row", Items: []models.SearchFilterSubItem{
		{Text: "🍦 Cold Now", Tooltip: "Filter by cold weather"},
		{Text: "🌤️ Mild now", Tooltip: "Filter by mild weather"},
		{Text: "☀️ Warm now", Tooltip: "Filter by warm weather"},
	}},
	{Type: "row", Items: []models.SearchFilterSubItem{
		{Text: "💵<US$1k/mo", Tooltip: "Filter by cost < 1k"},
		{Text: "💵<US$2k/mo", Tooltip: "Filter by cost < 2k"},
		{Text: "💵<US$3k/mo", Tooltip: "Filter by cost < 3k"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "👮‍ Safe", Tooltip: "Filter by safety"},
		{Text: "📡 Fast interne", Tooltip: "Filter by internet speed"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "💨 Clean air now", Tooltip: "Filter by air quality"},
		{Text: "👍 Liked by members", Tooltip: "Filter by member likes"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🔥 Popular now", Tooltip: "Filter by popularity"},
		{Text: "📈 Growing in nomads", Tooltip: "Filter by growth"},
	}},
	{Type: "row-two-stick", Items: []models.SearchFilterSubItem{
		{Text: "🏅 Top ranked", Tooltip: "Filter by ranking"},
		{Text: "💎 Hidden gem", Tooltip: "Filter by hidden gems"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "✨ You haven't been", Tooltip: "Places you haven't visited"},
		{Text: "✨ For you", Tooltip: "Recommendations for you"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🇸🇬 Near Singapore", Tooltip: "Filter by proximity to Singapore", Class: "two-btn-one-box"},
	}},
	{Type: "header", Text: "WHERE"},
	{Type: "where-grid", Items: []models.SearchFilterSubItem{
		{Text: "🌎North America", Tooltip: "🌎North America"},
		{Text: "💃Latin America", Tooltip: "💃Latin America"},
		{Text: "🇪🇺Europe", Tooltip: "🇪🇺Europe"},
		{Text: "🌎Africa", Tooltip: "🌎Africa"},
		{Text: "🕌 Middle East", Tooltip: "🕌 Middle East"},
		{Text: "⛩️ Asia", Tooltip: "⛩️ Asia"},
		{Text: "🏄Oceania", Tooltip: "Oceania"},
		{Text: "🛰️Space", Tooltip: "Space"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🇺🇸 United States", Tooltip: "🇺🇸 United States"},
		{Text: "🏝 Caribbean", Tooltip: "🏝 Caribbean"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🇪🇺 European Union", Tooltip: "European Union"},
		{Text: "🇪🇺 Not in Schengen", Tooltip: "Not in Schengen"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🤝 Easy to make friends", Tooltip: "Easy to make friends"},
		{Text: "❤️ Great for dating", Tooltip: "Great for dating"},
	}},
	{Type: "header", Text: "WHEN"},
	{Type: "row-two-stick", Items: []models.SearchFilterSubItem{
		{Text: "☃️ In the winter", Tooltip: "Winter"},
		{Text: "♻️ All year round", Tooltip: "All year"},
	}},
	{Type: "where-grid-when", MonthItems: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}},
	{Type: "header", Text: "RATING"},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "🤩 Exceptional (4.75+)", Tooltip: "Exceptional"},
		{Text: "😍 Very good (4.5+)", Tooltip: "Very good"},
	}},
	{Type: "row-two", Items: []models.SearchFilterSubItem{
		{Text: "👍 Good (4.25+)", Tooltip: "Good"},
		{Text: "😐 Okay (3+)", Tooltip: "Okay"},
	}},
	{Type: "header", Text: "VISA FREE FOR"},
	{Type: "select", ID: "select-pass-2", Placeholder: "select your passport", Options: []string{"🇦🇫 Afganistan"}},
	{Type: "plus"},
	{Type: "select", ID: "select-pass-3", Placeholder: "select your passport", Options: []string{"🇦🇫 Afganistan"}},
	{Type: "plus"},
	{Type: "select", ID: "select-pass", Placeholder: "select your passport", Options: []string{"🇦🇫 Afganistan"}},
}
