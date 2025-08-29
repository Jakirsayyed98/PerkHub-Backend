# Build stage
FROM golang:1.23 AS build

#FROM alpine:latest
#RUN apk add --no-cache bash curl

WORKDIR /src

ENV GOCACHE=/tmp/.cache

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
COPY migrations /app/migrations

# Build binary (output in /src)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-api .

# Optionally set default mode
ENV GIN_MODE=release

# Clean up Go build caches to reduce image size
RUN go clean -cache -modcache -testcache -fuzzcache

# Runtime stage
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /src/go-api /app/go-api

# Copy the .env file explicitly
COPY --from=build /src/.env.production /app/.env.production

USER nonroot:nonroot
EXPOSE 4215
ENTRYPOINT ["/app/go-api"]
