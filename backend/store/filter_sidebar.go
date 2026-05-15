package store

import "backend/models"

var FilterSidebar = []models.FilterSidebarSection{
	{
		Header: "WHAT",
		Rows: []models.FilterSidebarRow{
			{Classes: "row-filter", Items: []models.FilterSidebarItem{
				{Text: "🍦 Cold Now", Tooltip: "Filter by cold weather", Classes: "tooltip"},
				{Text: "🌤️ Mild now", Tooltip: "Filter by mild weather", Classes: "tooltip"},
				{Text: "☀️ Warm now", Tooltip: "Filter by warm weather", Classes: "tooltip"},
			}},
			{Classes: "row-filter", Items: []models.FilterSidebarItem{
				{Text: "💵<US$1k/mo", Tooltip: "Filter by cold weather", Classes: "tooltip"},
				{Text: "💵<US$2k/mo", Tooltip: "Filter by mild weather", Classes: "tooltip"},
				{Text: "💵<US$3k/mo", Tooltip: "Filter by warm weather", Classes: "tooltip"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "👮‍ Safe", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "📡 Fast interne", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "💨 Clean air now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "👍 Liked by members", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🔥 Popular now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "📈 Growing in nomads", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two stick", Items: []models.FilterSidebarItem{
				{Text: "🏅 Top ranked", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "💎 Hidden gem", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "✨ You haven't been", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "✨ For you", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🇸🇬 Near Singapore", Tooltip: "", Classes: "tooltip two-btn-one-box"},
			}},
		},
	},
	{
		Header: "WHERE",
		Grid: &models.FilterSidebarGrid{
			Classes: "where-grid",
			Items: []models.FilterSidebarGridItem{
				{Text: "🌎North America", Tooltip: "🌎North America", Element: "h3", Classes: "tooltip"},
				{Text: "💃Latin America ", Tooltip: "💃Latin America ", Element: "h3", Classes: "tooltip"},
				{Text: "🇪🇺Europe", Tooltip: "🇪🇺Europe", Element: "h3", Classes: "tooltip"},
				{Text: "🌎Africa", Tooltip: "🌎Africa", Element: "h3", Classes: "tooltip"},
				{Text: "🕌 Middle East", Tooltip: "🕌 Middle East", Element: "h3", Classes: "tooltip"},
				{Text: "⛩️ Asia", Tooltip: "⛩️ Asia", Element: "h3", Classes: "tooltip"},
				{Text: "🏄Oceania", Tooltip: "", Element: "h3", Classes: "tooltip"},
				{Text: "🛰️Space", Tooltip: "", Element: "h3", Classes: "tooltip"},
			},
		},
		Rows: []models.FilterSidebarRow{
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🇺🇸 United States", Tooltip: "🇺🇸 United States", Classes: "tooltip two-btn"},
				{Text: "🏝 Caribbean", Tooltip: "🏝 Caribbean", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🇪🇺 European Union", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🇪🇺 Not in Schengen", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🤝 Easy to make friends", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "❤️ Great for dating", Tooltip: "", Classes: "tooltip two-btn"},
			}},
		},
	},
	{
		Header: "WHEN",
		Grid: &models.FilterSidebarGrid{
			Classes: "where-grid when",
			Items: []models.FilterSidebarGridItem{
				{Text: "Jan", Element: "h3"},
				{Text: "Feb", Element: "h3"},
				{Text: "Mar", Element: "h3"},
				{Text: "Apr", Element: "h3"},
				{Text: "May", Element: "h3"},
				{Text: "Jun", Element: "h3"},
				{Text: "Jul", Element: "h3"},
				{Text: "Aug", Element: "h3"},
				{Text: "Sep", Element: "h3"},
				{Text: "Oct", Element: "h3"},
				{Text: "Nov", Element: "h3"},
				{Text: "Dec", Element: "h3"},
			},
		},
		Rows: []models.FilterSidebarRow{
			{Classes: "row-filter two stick", Items: []models.FilterSidebarItem{
				{Text: "☃️ In the winter", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "♻️ All year round", Tooltip: "", Classes: "tooltip two-btn"},
			}},
		},
	},
	{
		Header: "WHEN",
		Rows: []models.FilterSidebarRow{
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🤩 Exceptional (4.75+)", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: " 😍 Very good (4.5+)", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "👍 Good (4.25+)", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "😐 Okay (3+)", Tooltip: "", Classes: "tooltip two-btn"},
			}},
		},
	},
	{
		Header: "VISA FREE FOR",
		Rows: []models.FilterSidebarRow{
			{Type: "select", ID: "select-pass-2", Classes: "row-filter passport", Options: []string{"Afganistan", "Albania", "Algeria", "Andorra", "Angola"}},
			{Type: "divider", Text: "+"},
			{Type: "select", ID: "select-pass-3", Classes: "row-filter passport", Options: []string{"Afganistan", "Albania", "Algeria", "Andorra", "Angola"}},
			{Type: "divider", Text: "+"},
			{Type: "select", ID: "select-pass", Classes: "row-filter passport", Options: []string{"Afganistan", "Albania", "Algeria", "Andorra", "Angola"}},
		},
	},
	{
		Header: "OTHER TIMES",
		Rows: []models.FilterSidebarRow{
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
			{Classes: "row-filter two", Items: []models.FilterSidebarItem{
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
				{Text: "🌤️ Mild now", Tooltip: "", Classes: "tooltip two-btn"},
			}},
		},
	},
}
