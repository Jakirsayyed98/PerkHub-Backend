# Dockerfile
# Build stage
FROM golang:1.23 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/go-api ./cmd/main  # adjust package path

# Runtime stage (small)
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /app/go-api /app/go-api
# create non-root user at runtime
USER nonroot:nonroot
EXPOSE 8081
ENTRYPOINT ["/app/go-api"]
