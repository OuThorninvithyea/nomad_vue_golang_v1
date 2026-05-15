package models

type NewMembersCard struct {
	Header      string   `json:"header"`
	PrivacyText string   `json:"privacyText"`
	Members     []Member `json:"members"`
}

type Member struct {
	ID    int    `json:"id"`
	Thumb string `json:"thumb"`
	Blur  string `json:"blur"`
}
