# Build stage
FROM golang:1.23 AS build
WORKDIR /src

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (output in /src)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-api .

# Runtime stage
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /src/go-api /app/go-api

USER nonroot:nonroot
EXPOSE 8081
ENTRYPOINT ["/app/go-api"]
