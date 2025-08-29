# Build stage
FROM golang:1.24-alpine AS builder

# Install git and build tools
RUN apk add --no-cache git build-base
ENV GOCACHE=/tmp/.cache

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Set working directory
WORKDIR /app

# Copy the entire repository first (including .git directory)
COPY . .

# Print directory structure to debug
RUN find . -type d | sort

# Optionally set default mode
ENV GIN_MODE=release

# Clean up Go build caches to reduce image size
RUN go clean -cache -modcache -testcache -fuzzcache

# Fetch dependencies
RUN go mod download

# Build the application with optimizations
RUN GIN_MODE=release GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags=jsoniter -ldflags="-w -s" -o go-api

# Runtime stage - using Alpine image as a base
FROM alpine:3.18

# Install bash and other necessary utilities for debugging and interaction
RUN apk --no-cache add bash ca-certificates tzdata fontconfig

# Set timezone to Asia/Kolkata (IST)
ENV TZ=Asia/Kolkata

# Create app directory
WORKDIR /app

# Copy the go-api binary from the build stage
COPY --from=builder /app/go-api .

# Copy migrations the build stage
COPY --from=builder /app/.env.production /app/.env.production
COPY --from=builder /app/migrations /app/migrations

# Create logs directory
RUN mkdir -p /app/logs && chmod 755 /app/logs

# Print directory structure to verify
RUN echo "Final image structure:" && find /app -type d | sort

# Environment variables
ENV GIN_MODE=release
ENV AWS_REGION=ap-south-1
ENV AWS_DEFAULT_REGION=ap-south-1

# Expose the application port
EXPOSE 4215

# Command to run the application
CMD ["./go-api"]
