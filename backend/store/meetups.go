package store

import "backend/models"

var Meetups = models.MeetupsCard{
	CardType:   "meetups",
	Header:     "🥥 Next meetups (15/mo)",
	FooterText: "See upcoming meetups →",
	FooterHref: "#",
	Days: []models.MeetupDay{
		{
			CityImage:    "/images/asuncion-paraguay.webp",
			CityImageAlt: "Asuncion",
			Schedule:     "Mon 9th Mar: Barcelona",
			Rsvps:        "1 RSVPS",
			AttendeeImages: []string{
				"/images/profile-13.jpg",
				"/images/profile-7.jpg",
			},
		},
		{
			CityImage:    "/images/barcelona-spain.webp",
			CityImageAlt: "Barcelona",
			Schedule:     "Thu 5th Mar: Asuncion",
			Rsvps:        "3 RSVPS",
			AttendeeImages: []string{
				"/images/profile-15.jpg",
				"/images/profile-4.jpg",
				"/images/profile-10.jpg",
				"/images/profile-1.jpg",
				"/images/profile-8.jpg",
				"/images/profile-9.jpg",
				"/images/profile-4.webp",
				"/images/profile-14.webp",
			},
		},
	},
}
