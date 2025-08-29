# Build stage
FROM golang:1.23 AS build

WORKDIR /src

ENV GOCACHE=/tmp/.cache

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Install bash for debugging (needed for the build stage)
RUN apt-get update && apt-get install -y bash

# Copy source code
COPY . .
COPY --from=build /src/migrations /app/migrations

# Build binary (output in /src)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-api .

# Optionally set default mode
ENV GIN_MODE=release

# Clean up Go build caches to reduce image size
RUN go clean -cache -modcache -testcache -fuzzcache

# Debug stage (temporary for debugging purposes)
FROM debian:bullseye-slim  # Use a minimal Debian image with bash support

# Install bash and any other debugging tools
RUN apt-get update && apt-get install -y bash curl

WORKDIR /app

# Copy build artifacts from the build stage
COPY --from=build /src/go-api /app/go-api
COPY --from=build /src/migrations /app/migrations
COPY --from=build /src/.env.production /app/.env.production

# Expose the port (for API)
EXPOSE 4215

# Switch to root user for debugging
USER root

# The command to run the application
ENTRYPOINT ["/app/go-api"]
