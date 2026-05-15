package store

import "backend/models"

var Hero = models.HeroData{
	BackgroundVideo: models.HeroVideo{
		Src:  "/images/nomad-vdo-hd.mp4",
		Type: "video/mp4",
	},
	Left: models.HeroLeft{
		LaurelImage: "/images/laurel.svg",
		AwardHeader: "#1 Nomad Community",
		AwardFooter: "Since 2014",
		Stars:       5,
		MainTitle:   "Go nomad",
		Emoji:       "🌍",
		SubTitle:    "Join a global community of remote workers living and traveling around the world",
		RemoteTravelers: []string{
			"/images/profile-1.jpg",
			"/images/profile-2.jpg",
			"/images/profile-3.jpg",
			"/images/profile-4.webp",
			"/images/profile-5.jpg",
			"/images/profile-6.jpg",
			"/images/profile-7.webp",
			"/images/profile-8.jpg",
			"/images/profile-9.webp",
			"/images/profile-10.jpg",
			"/images/profile-11.jpg",
			"/images/profile-12.webp",
		},
		Benefits: []models.HeroBenefit{
			{Emoji: "🍹", Text: "Attend 283 meetups/year", Context: "in 100+ cities", Href: "#"},
			{Emoji: "❤️", Text: "Meet new people", Context: "for dating and friends", Href: "#"},
			{Emoji: "🔬", Text: "Research destinations", Context: "and find your best place to live and work", Href: "#"},
			{Emoji: "🌍", Text: "Keep track of your travels", Context: "and record where you've been", Href: "#"},
			{Emoji: "💬", Text: "Join Nomads.com chat", Context: "and find your community on the road", Href: "#"},
		},
	},
	Right: models.HeroRight{
		Video:            models.HeroMedia{Src: "/images/video1-thumb.jpg", Alt: "video-thumbnail"},
		PlayButton:       models.HeroMedia{Src: "/images/pause.svg", Alt: "play-button"},
		InputPlaceholder: "Type your email...",
		CtaHref:          "#join.com",
		CtaText:          "Join Nomads.com →",
		LoginText:        "if you already have an account, we'll log you in",
	},
}
