package store

import "backend/models"

var BoxHover = models.BoxHoverData{
	HoverHeader: models.BoxHoverHeader{
		CloseTooltip: "Hide this, it will come back if you filter or reload Nomads.com",
		CloseLabel:   "×",
	},
}
