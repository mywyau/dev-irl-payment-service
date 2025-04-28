# Use a lightweight official Go image
FROM golang:1.22-alpine as builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app
RUN go build -o payment-service ./cmd/server

# Final lightweight image
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/payment-service .

# Set environment variables if needed
ENV PORT=8080

# Expose port
EXPOSE 8080

# Command to run
CMD ["./payment-service"]
