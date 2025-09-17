# ---- Build stage ----
FROM golang:1.24.6 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary from root main.go
RUN go build -o /app/movie-reservation ./main.go

# ---- Run stage ----
FROM debian:bullseye-slim

WORKDIR /app

# Install CA certificates for HTTPS if needed
# RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary from builder
COPY --from=builder /app/movie-reservation /app/movie-reservation

# Run the binary
CMD ["/app/movie-reservation"]
