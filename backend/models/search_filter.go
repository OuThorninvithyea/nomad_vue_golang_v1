package models

type SearchFilterItem struct {
	Type        string                `json:"type"`
	Text        string                `json:"text,omitempty"`
	Items       []SearchFilterSubItem `json:"items,omitempty"`
	MonthItems  []string              `json:"monthItems,omitempty"`
	ID          string                `json:"id,omitempty"`
	Placeholder string                `json:"placeholder,omitempty"`
	Options     []string              `json:"options,omitempty"`
}

type SearchFilterSubItem struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class,omitempty"`
}
