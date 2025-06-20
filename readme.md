
# FlagFairway

FlagFairway is a lightweight, self-hosted feature flag service.  
It provides a simple REST API for managing feature flags, a modern Preact-based web UI, and persistent storage using BadgerDB.

---


## what this is intended for

- having feature flags (server side) for backend applications

## What this is lacking at the momment

- Security (please keep this behind thne firewall or add authentication)
- scaleablity (this is designed for small to medium size projects at the moment)

## Features

- 🚩 **Feature Flag API**: Create, update, delete, and list feature flags via REST endpoints.
- 💾 **Persistent Storage**: Uses [BadgerDB](https://github.com/dgraph-io/badger) for fast, embedded key-value storage.
- 🕒 **Automatic DB Maintenance**: Scheduled garbage collection for optimal performance.
- 🌐 **Modern Web UI**: Preact + TypeScript + Sass frontend for easy flag management.
- ⚡ **Fast & Lightweight**: Minimal dependencies, quick startup, and efficient resource usage.

---

## Project Structure

```
FlagFairway/
│
├── main.go                # Go backend (API server)
├── internal/
│   └── ui/
│       ├── src/           # Frontend source (Preact, TSX, Sass)
│       ├── static/        # Static assets (favicon, index.html, etc.)
│       ├── build/         # Compiled frontend output
│       ├── build.js       # esbuild build script
│       ├── package.json   # Frontend dependencies/scripts
│       ├── tsconfig.json  # TypeScript config
│       └── README.md      # UI-specific documentation
├── data/                  # BadgerDB data directory (created at runtime)
└── README.md              # (this file)
```

---

## Getting Started

### 1. Backend (Go API)

#### Prerequisites

- Go 1.18+
- [BadgerDB](https://github.com/dgraph-io/badger) (installed via Go modules)

#### Run the server

```sh
go run main.go
```

- The server listens on `http://localhost:8080`
- The web UI is served from `/` (after building the frontend)
- API endpoints are under `/api/v1/feature_flags`

#### API Endpoints

| Method | Endpoint                                 | Description                      |
|--------|------------------------------------------|----------------------------------|
| GET    | `/api/v1/feature_flags`                  | List all feature flag keys       |
| GET    | `/api/v1/feature_flags/:key`             | Get value for a feature flag     |
| POST   | `/api/v1/feature_flags/:key`             | Create a new feature flag        |
| PUT    | `/api/v1/feature_flags/:key`             | Update an existing feature flag  |
| DELETE | `/api/v1/feature_flags/:key`             | Delete a feature flag            |
| GET    | `/api/v1/health`                         | Health check                     |

---

### 2. Frontend (UI)

See [`internal/ui/README.md`](internal/ui/README.md) for full details.

#### Quick Start

```sh
cd internal/ui
npm install
npm run build
```

- The build output will be in `internal/ui/build/`
- The Go server will serve this UI automatically

#### Development

- Edit UI code in `internal/ui/src/`
- Styles are in `internal/ui/src/style.sass`
- Run `npm run build` after changes

---

## Deployment

1. Build the frontend (`npm run build` in `internal/ui`)
2. Run the Go server (`go run main.go`)
3. Visit [http://localhost:8080](http://localhost:8080) in your browser

---

## Customization

- **UI**: Modify or extend the Preact components in `internal/ui/src/`
- **API**: Extend Go handlers in `main.go`
- **Storage**: BadgerDB data is stored in the `data/` directory by default

---

## License

MIT

---

## Credits

- [Preact](https://preactjs.com/)
- [BadgerDB](https://github.com/dgraph-io/badger)
- [Echo Web Framework](https://echo.labstack.com/)
- [esbuild](https://esbuild.github.io/)
