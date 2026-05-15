package models

type HeroData struct {
	BackgroundVideo HeroVideo `json:"backgroundVideo"`
	Left            HeroLeft  `json:"left"`
	Right           HeroRight `json:"right"`
}

type HeroVideo struct {
	Src  string `json:"src"`
	Type string `json:"type"`
}

type HeroLeft struct {
	LaurelImage     string        `json:"laurelImage"`
	AwardHeader     string        `json:"awardHeader"`
	AwardFooter     string        `json:"awardFooter"`
	Stars           int           `json:"stars"`
	MainTitle       string        `json:"mainTitle"`
	Emoji           string        `json:"emoji"`
	SubTitle        string        `json:"subTitle"`
	RemoteTravelers []string      `json:"remoteTravelers"`
	Benefits        []HeroBenefit `json:"benefits"`
}

type HeroBenefit struct {
	Emoji   string `json:"emoji"`
	Text    string `json:"text"`
	Context string `json:"context"`
	Href    string `json:"href"`
}

type HeroRight struct {
	Video            HeroMedia `json:"video"`
	PlayButton       HeroMedia `json:"playButton"`
	InputPlaceholder string    `json:"inputPlaceholder"`
	CtaHref          string    `json:"ctaHref"`
	CtaText          string    `json:"ctaText"`
	LoginText        string    `json:"loginText"`
}

type HeroMedia struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}
