# Use buildkit
# syntax=docker/dockerfile:experimental

FROM node:bookworm-slim as uibuilder
WORKDIR /app
COPY ./internal/ui .
RUN npm install -g vite && npm install
RUN --mount=type=cache,target=/root/.npm npm run build


FROM golang:bookworm AS serverbuilder
ENV GOPROXY=${GOPROXY}
WORKDIR /app
# COPY go.mod ./
COPY . .
RUN go mod download
RUN --mount=type=cache,target=/go/pkg/mod CGO_ENABLED=0 go build -o app

# Final stage
FROM gcr.io/distroless/static-debian12
WORKDIR /app
# Copy ui
COPY --from=uibuilder /app/build ./internal/ui/build

# Copy server
COPY --from=serverbuilder /app/app ./

#Volume setup for data directory
VOLUME [ "/data" ]

EXPOSE 8080
CMD ["./app"]