# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:latest@sha256:5f3787b7f902c07c7ec4f3aa91a301a3eda8133aa32661a3b3a3a86ab3a68a36 AS serverbuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app


FROM node:lts-alpine@sha256:01743339035a5c3c11a373cd7c83aeab6ed1457b55da6a69e014a95ac4e4700b AS uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN  --mount=type=cache,target=/root/.npm npm ci --ignore-scripts --prefer-offline
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


# Final stage
FROM gcr.io/distroless/static-debian13@sha256:47b2d72ff90843eb8a768b5c2f89b40741843b639d065b9b937b07cd59b479c6
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
