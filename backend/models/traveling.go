package models

type TravelingCard struct {
	Header      string   `json:"header"`
	PrivacyText string   `json:"privacyText"`
	Members     []Member `json:"members"`
}
