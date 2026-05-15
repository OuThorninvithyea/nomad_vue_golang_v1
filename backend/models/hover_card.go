package models

type HoverCard struct {
	Quote          string `json:"quote"`
	Image          string `json:"image"`
	ImageAlt       string `json:"imageAlt"`
	AnimationClass string `json:"animationClass"`
}
