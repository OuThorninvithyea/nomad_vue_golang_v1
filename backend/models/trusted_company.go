package models

type TrustedCompany struct {
	Header    string    `json:"header"`
	Companies []Company `json:"companies"`
}

type Company struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}
