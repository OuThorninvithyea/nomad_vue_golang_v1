package models

type ChatCard struct {
	Header   string `json:"header"`
	Image    string `json:"image"`
	ImageAlt string `json:"imageAlt"`
	CtaText  string `json:"ctaText"`
	CtaHref  string `json:"ctaHref"`
}
