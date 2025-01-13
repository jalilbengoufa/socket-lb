# server/Dockerfile

# Stage 1: Build the Go binary
FROM golang:1.21.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o websocket-server main.go

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/websocket-server .

# Expose the port the server listens on
EXPOSE 8080

# Command to run the executable
CMD ["./websocket-server"]

