package models

type BoxHoverData struct {
	HoverHeader BoxHoverHeader `json:"hoverHeader"`
}

type BoxHoverHeader struct {
	CloseTooltip string `json:"closeTooltip"`
	CloseLabel   string `json:"closeLabel"`
}
