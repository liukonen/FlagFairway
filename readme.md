# FlagFairway
Tags: `Go` `BadgerDB` `Preact` `Feature Flags` `REST API` `Sass`

A lightweight, self-hosted feature flag service with a modern web UI and persistent storage.

## Context & Story
FlagFairway was built to provide a simple, efficient solution for managing feature flags in backend applications. It is designed for small to medium-sized projects, prioritizing ease of use and minimal dependencies. While it lacks advanced scalability and security features, it serves as a robust starting point for teams needing server-side feature flagging.

## Architecture & Decisions
- **Go Backend**: Chosen for its performance and simplicity in building REST APIs.
- **BadgerDB**: Embedded key-value storage for fast, persistent data handling.
- **Preact Frontend**: Lightweight alternative to React for a modern UI.
- **Sass**: Used for styling to maintain modular and maintainable CSS.
- **Echo Framework**: Selected for its simplicity and middleware support.
- **Cron Jobs**: Automated database garbage collection for optimal performance.

## Key Features
- **Feature Flag API**: REST endpoints to create, update, delete, and list feature flags.
- **Persistent Storage**: Utilizes BadgerDB for efficient data storage.
- **Automatic Maintenance**: Scheduled garbage collection to optimize database performance.
- **Modern Web UI**: Built with Preact, TypeScript, and Sass for intuitive flag management.
- **Lightweight Design**: Minimal dependencies and efficient resource usage.

## Quick Start
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/liukonen/FlagFairway.git
   cd FlagFairway
   ```

2. **Run the Backend**:
   Ensure you have Go installed.
   ```bash
   go run main.go
   ```

3. **Build the Frontend**:
   Navigate to the `internal/ui` directory and install dependencies:
   ```bash
   cd internal/ui
   npm install
   npm run build
   ```

4. **Access the Application**:
   Open your browser and navigate to `http://localhost:8080`.

5. **API Endpoints**:
   - `GET /api/v1/feature_flags`: List all feature flags.
   - `POST /api/v1/feature_flags/:key`: Create a new feature flag.
   - `PUT /api/v1/feature_flags/:key`: Update an existing feature flag.
   - `DELETE /api/v1/feature_flags/:key`: Delete a feature flag.
   - `GET /api/v1/health`: Check service health.

---

### Notes
- **Security**: Keep the service behind a firewall or add authentication.
- **Scalability**: Designed for small to medium projects; not optimized for high-scale environments.
