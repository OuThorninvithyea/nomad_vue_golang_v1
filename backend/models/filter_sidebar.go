package models

type FilterSidebarSection struct {
	Header string                `json:"header"`
	Rows   []FilterSidebarRow    `json:"rows"`
	Grid   *FilterSidebarGrid    `json:"grid,omitempty"`
}

type FilterSidebarRow struct {
	Classes string               `json:"classes"`
	Type    string               `json:"type,omitempty"`
	ID      string               `json:"id,omitempty"`
	Text    string               `json:"text,omitempty"`
	Items   []FilterSidebarItem  `json:"items,omitempty"`
	Options []string             `json:"options,omitempty"`
}

type FilterSidebarItem struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Classes string `json:"classes"`
}

type FilterSidebarGrid struct {
	Classes string                  `json:"classes"`
	Items   []FilterSidebarGridItem `json:"items"`
}

type FilterSidebarGridItem struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip,omitempty"`
	Element string `json:"element"`
	Classes string `json:"classes,omitempty"`
}
