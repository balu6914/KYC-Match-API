# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kyc-match-api main.go

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests (HarperDB)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/kyc-match-api .

# Copy .env file (optional, for local testing)
COPY .env .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./kyc-match-api"]