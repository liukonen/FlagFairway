FROM node:bookworm-slim as uibuilder
WORKDIR /app
COPY ./internal/ui/package.json ./
COPY ./internal/ui/package-lock.json ./
RUN npm install -g vite && npm install
COPY ./internal/ui .
RUN npm run build


FROM golang:bookworm AS serverbuilder
ENV GOPROXY=${GOPROXY}
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

# Final stage
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=serverbuilder /app/app ./
COPY --from=serverbuilder /app/docs/swagger.json ./docs/swagger.json
COPY --from=uibuilder /app/build ./internal/ui/build

#Volume setup for data directory
VOLUME [ "/data" ]
EXPOSE 8080
CMD ["./app"]