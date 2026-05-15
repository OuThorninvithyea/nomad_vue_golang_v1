package routes


import (
	"backend/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/search-dropdown", handlers.SearchDropdown)
	app.Get("/filter-sidebar", handlers.FilterSidebar)
	app.Get("/chat", handlers.Chat)
	app.Get("/new-members", handlers.NewMembers)
	app.Get("/box-hover", handlers.BoxHover)
	app.Get("/today-pick", handlers.TodayPick)
	app.Get("/traveling", handlers.Traveling)
	app.Get("/trusted-company", handlers.TrustedCompany)
	app.Get("/cities", handlers.Cities)
	app.Get("/hover-card", handlers.HoverCard)
	app.Get("/ad-one", handlers.AdOne)
	app.Get("/ad-two", handlers.AdTwo)
	app.Get("/suggest", handlers.Suggest)
	app.Get("/hero", handlers.Hero)
	app.Get("/logo-dropdown", handlers.LogoDropdown)
	app.Get("/meetups", handlers.Meetups)
	app.Get("/search-filter", handlers.SearchFilter)
	app.Get("/cities", handlers.Cities)
}