# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:latest@sha256:e2ddb153f786ee6210bf8c40f7f35490b3ff7d38be70d1a0d358ba64225f6428 AS serverbuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app


FROM node:lts-alpine@sha256:7fddd9ddeae8196abf4a3ef2de34e11f7b1a722119f91f28ddf1e99dcafdf114 AS uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN  --mount=type=cache,target=/root/.npm npm ci --ignore-scripts --prefer-offline
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


# Final stage
FROM gcr.io/distroless/static-debian13@sha256:28efbe90d0b2f2a3ee465cc5b44f3f2cf5533514cf4d51447a977a5dc8e526d0
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
