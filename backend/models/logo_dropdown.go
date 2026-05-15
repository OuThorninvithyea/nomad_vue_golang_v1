package models

type LogoDropdownSection struct {
	SectionClass string             `json:"sectionClass,omitempty"`
	Title        string             `json:"title,omitempty"`
	Items        []LogoDropdownItem `json:"items"`
}

type LogoDropdownItem struct {
	Text    string `json:"text"`
	Emoji   string `json:"emoji"`
	Badge   string `json:"badge,omitempty"`
	Tooltip string `json:"tooltip,omitempty"`
}
