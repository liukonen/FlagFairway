# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:1.27.3-bookworm@sha256:dc0963f32053e4bc451d24d87e3892d42d01021b61f093e98eb95c0352bb84b0 AS serverbuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app


FROM node:26-slim@sha256:e6bedef217677b276c591795e169840edd9096dbf13f29df8ef8d1856b8820bc AS uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN  --mount=type=cache,target=/root/.npm npm ci --ignore-scripts --prefer-offline
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


# Final stage
FROM gcr.io/distroless/static-debian13@sha256:448810357cc8e67b7208988abdba0f5ec3d3aebee0a5b13eb624979792cffc0f
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
