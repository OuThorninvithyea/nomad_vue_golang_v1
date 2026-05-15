package models

type MeetupsCard struct {
	CardType   string      `json:"cardType"`
	Header     string      `json:"header"`
	FooterText string      `json:"footerText"`
	FooterHref string      `json:"footerHref"`
	Days       []MeetupDay `json:"days"`
}

type MeetupDay struct {
	CityImage      string   `json:"cityImage"`
	CityImageAlt   string   `json:"cityImageAlt"`
	Schedule       string   `json:"schedule"`
	Rsvps          string   `json:"rsvps"`
	AttendeeImages []string `json:"attendeeImages"`
}
