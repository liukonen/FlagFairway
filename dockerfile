# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:bookworm AS serverbuilder
ENV GOPROXY=${GOPROXY}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app


FROM node:lts-alpine as uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN npm ci
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


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