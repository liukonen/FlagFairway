# Use buildkit
# syntax=docker/dockerfile:experimental

FROM golang:latest@sha256:06d1251c59a75761ce4ebc8b299030576233d7437c886a68b43464bad62d4bb1 AS serverbuilder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app


FROM node:lts-alpine@sha256:cd6fb7efa6490f039f3471a189214d5f548c11df1ff9e5b181aa49e22c14383e AS uibuilder
WORKDIR /app
COPY ./internal/ui/package*.json ./
RUN  --mount=type=cache,target=/root/.npm npm ci --ignore-scripts --prefer-offline
COPY ./internal/ui .
RUN npm run build
RUN rm -rf /app/node_modules /app/.npm


# Final stage
FROM gcr.io/distroless/static-debian13@sha256:972618ca78034aaddc55864342014a96b85108c607372f7cbd0dbd1361f1d841
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
