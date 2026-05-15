package models

type AdOne struct{
	Image    string `json:"image"`
	ImageAlt string `json:"imageAlt"`
	Heading  string `json:"heading"`
	CtaText  string `json:"ctaText"`
	CtaHref  string `json:"ctaHref"`
	AdLabel  string `json:"adLabel"`

	
}
