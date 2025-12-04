# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:1.25.5-bookworm AS serverbuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app


FROM node:24.11.1-alpine3.21 as uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN npm ci
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


# Final stage
FROM gcr.io/distroless/static-debian12@sha256:4b2a093ef4649bccd586625090a3c668b254cfe180dee54f4c94f3e9bd7e381e
WORKDIR /app
# Copy ui
COPY --from=uibuilder /app/build ./internal/ui/build
# Copy server
COPY --from=serverbuilder /app/app ./

#Volume setup for data directory
VOLUME [ "/data" ]

EXPOSE 8080
USER 1000:1000
ENTRYPOINT ["/app/app"]
