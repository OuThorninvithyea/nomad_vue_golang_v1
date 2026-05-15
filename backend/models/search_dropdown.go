package models

type SearchDropdownGroup struct {
	ID    string               `json:"id"`
	Items []SearchDropdownItem `json:"items"`
}

type SearchDropdownItem struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
}
