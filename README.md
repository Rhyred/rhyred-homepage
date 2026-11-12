# RHYRED.NET

> Personal Developer Portfolio & Digital Laboratory.
> Built with Go. Powered by curiosity.

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=flat&logo=go)
![Templ](https://img.shields.io/badge/Templ-0.3-purple?style=flat)
![HTMX](https://img.shields.io/badge/HTMX-2.0-blue?style=flat)

---

## Overview

RHYRED.NET is a personal developer homepage built with **Go**, **Templ**, **HTMX**, and **Vanilla CSS**. It represents a developer exploring software engineering, artificial intelligence, networking, IoT, and computer vision.

### Design Direction

- Hybrid Editorial × Developer Terminal aesthetic
- Dark editorial layout with orange accent
- Large typography with asymmetric grid
- Responsive across desktop, tablet, and mobile

---

## Technology Stack

| Layer      | Technology                          |
|------------|-------------------------------------|
| Backend    | Go (standard library `net/http`)    |
| Templates  | Templ                               |
| Interaction| HTMX                                |
| Styling    | Vanilla CSS                         |
| JavaScript | Vanilla JS (minimal)                |
| Data       | JSON files                          |

---

## Project Structure

```
rhyred-homepage/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # App configuration
│   ├── handlers/                # HTTP handlers + tests
│   ├── models/                  # Data models
│   ├── router/                  # HTTP router setup
│   └── services/                # Data loading service
├── web/
│   ├── templates/
│   │   ├── layouts/             # Base HTML layout (Templ)
│   │   ├── components/          # Reusable UI components (Templ)
│   │   └── pages/               # Page templates (Templ)
│   ├── static/
│   │   ├── css/                 # Modular CSS
│   │   ├── js/                  # Vanilla JavaScript
│   │   └── images/              # Static images
│   └── data/
│       ├── projects.json        # Project data
│       └── profile.json         # Profile & site data
├── .env.example
├── .gitignore
├── Dockerfile
├── Makefile
├── go.mod
└── go.sum
```

---

## Installation

### Prerequisites

- **Go** 1.27+ ([golang.org](https://golang.org))
- **Templ** CLI

### Install Templ

```powershell
go install github.com/a-h/templ/cmd/templ@latest
```

Ensure `templ` is available in your PATH (typically `$GOPATH/bin`).

### Clone & Setup

```powershell
git clone https://github.com/Rhyred/rhyred-homepage.git
cd rhyred-homepage
go mod tidy
```

---

## Development

### Generate Templates

```powershell
templ generate
```

### Run Development Server

```powershell
go run ./cmd/server
```

The server starts at **http://localhost:8080**.

### Run Tests

```powershell
go test ./... -v
```

### Format Code

```powershell
gofmt -w .
templ fmt .
```

---

## Build

### Build Binary

```powershell
go build -o bin/server.exe ./cmd/server
```

### Run Binary

```powershell
./bin/server.exe
```

### Docker

```powershell
docker build -t rhyred-homepage .
docker run -p 8080:8080 rhyred-homepage
```

---

## Routes

| Method | Path              | Description             |
|--------|-------------------|-------------------------|
| GET    | `/`               | Homepage                |
| GET    | `/about`          | About page              |
| GET    | `/projects`       | Projects listing        |
| GET    | `/projects/{slug}`| Project detail          |
| GET    | `/lab`            | Digital Laboratory      |
| GET    | `/journal`        | Journal                 |
| GET    | `/health`         | Health check (JSON)     |
| GET    | `/api/status`     | App status (JSON)       |
| GET    | `/static/*`       | Static assets           |

---

## Configuration

Configuration via environment variables (see `.env.example`):

| Variable   | Default          | Description         |
|------------|------------------|---------------------|
| `PORT`     | `8080`           | Server port         |
| `APP_NAME` | `RHYRED`         | Application name    |
| `BASE_URL` | `https://rhyred.net` | Base URL        |
| `DATA_DIR` | `web/data`       | Data directory path |

---

## Customization

- **Profile data**: Edit `web/data/profile.json`
- **Projects**: Edit `web/data/projects.json`
- **Styling**: Modify files in `web/static/css/`
- **Templates**: Modify `.templ` files and run `templ generate`

---

## License

© 2026 RHYRED. All rights reserved.

