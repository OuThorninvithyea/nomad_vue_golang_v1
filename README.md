<div align="center">
  <img src="./frontend/public/images/logo-trans.png" alt="Nomad List Logo" width="120" />
  <h1>Nomad Vue + Go</h1>
  <p><strong>Full-Stack Digital Nomad City Guide</strong></p>

  <p>
    <a href="https://cv-ou-thorninvithyea-fvkg.vercel.app/">
      <img src="https://img.shields.io/badge/Portfolio-Ou%20Thorninvithyea-blue?style=flat-square" alt="Portfolio" />
    </a>
    <img src="https://img.shields.io/badge/Vue-3.5-4FC08D?style=flat-square&logo=vue.js" alt="Vue 3" />
    <img src="https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go" alt="Go" />
    <img src="https://img.shields.io/badge/Vite-8.0-646CFF?style=flat-square&logo=vite" alt="Vite 8" />
    <img src="https://img.shields.io/badge/Status-Development-yellow?style=flat-square" alt="Status" />
  </p>
</div>

---

## Overview

**Nomad Vue + Go** is a full-stack digital nomad city guide platform that helps remote workers discover, compare, and connect with cities around the world. The Vue 3 frontend delivers a rich, interactive UI while the Go backend serves structured API data for cities, meetups, community members, and travel resources.

This platform combines visual storytelling with practical data to help nomads make informed decisions about where to live and work remotely.

## Key Features

### Frontend
- **City Database** — Explore 100+ cities with photos, flags, and key metrics
- **Smart Filtering** — Filter by region, cost of living, internet quality, activities
- **Community Hub** — Meetups, member profiles, chat cards, and testimonials
- **Hero Section** — Video background and engaging call-to-action
- **Nomad Score** — City comparison system
- **Trusted Companies** — Brand partner showcase

### Backend (Go)
- **RESTful API** — Gin framework with clean route separation
- **Data Stores** — In-memory stores for cities, members, testimonials, and more
- **Hot Reload** — Air configuration for development efficiency
- **Scalable Architecture** — Handler / Model / Store / Route pattern

## Tech Stack

| Layer | Technology |
|-------|-----------|
| **Frontend** | Vue 3 (Composition API) + Vite 8 + JavaScript |
| **Backend** | Go 1.23 + Gin |
| **Architecture** | RESTful API with in-memory data stores |
| **Dev Tools** | Air (hot reload for Go) |

## Project Structure

```
nomad_vue_golang_v1/
├── frontend/                    # Vue 3 SPA
│   ├── src/
│   │   ├── components/          # UI components
│   │   │   ├── cards/           # CityCard, ChatCard, MeetupsCard, etc.
│   │   │   ├── filterBar/       # Search, filters, grid controls
│   │   │   ├── hero/            # Hero section with video
│   │   │   ├── footer/          # Footer bar
│   │   │   ├── popup/           # Advertisement popup
│   │   │   ├── sidebar/         # Filter sidebar
│   │   │   ├── testimonials/    # User testimonials
│   │   │   └── trustedCompanies/ # Brand logos
│   │   ├── composables/         # API data hooks
│   │   ├── css/                 # Global styles
│   │   ├── views/               # Page views
│   │   └── main.js
│   └── ...
│
└── backend/                     # Go API server
    ├── main.go                  # Server entry point
    ├── go.mod / go.sum          # Dependencies
    ├── .air.toml                # Hot reload config
    ├── handlers/                # HTTP request handlers
    ├── models/                  # Data structures
    ├── store/                   # In-memory data stores
    └── routes/                  # Route registration
```

## Getting Started

### Prerequisites

- Node.js >= 20.19.0 or >= 22.12.0
- Go 1.23+
- Air (optional, for hot reload)

### Backend Setup

```bash
cd backend

# Run with hot reload (requires Air)
air

# Or run directly
go run main.go

# API server: http://localhost:8080
```

### Frontend Setup

```bash
cd frontend
npm install
npm run dev

# Dev server: http://localhost:5173
```

### Production Build

```bash
cd frontend
npm run build
```

## Author

**Ou Thorninvithyea**

- 🌐 [Portfolio](https://cv-ou-thorninvithyea-fvkg.vercel.app/)
- 🐙 [GitHub](https://github.com/OuThorninvithyea)
- 📧 Vithyeasa@gmail.com
- 📍 Phnom Penh, Cambodia

> Passionate Software Engineer with 2+ years of experience in full-stack web development, specializing in designing and maintaining scalable web applications.

---

<div align="center">
  <sub>Built with ❤️ by Ou Thorninvithyea</sub>
</div>
