# Step 1: Build the Go binary
FROM golang:1.24.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build fully static binary
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o movie-reservation main.go

# Step 2: Use scratch
FROM scratch

WORKDIR /app

COPY --from=builder /app/movie-reservation .

EXPOSE 8080

CMD ["./movie-reservation"]