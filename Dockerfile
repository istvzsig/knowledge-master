# Dockerfile
FROM golang:1.23.4 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o knowledge-manager

# Create a minimal image using distroless
FROM gcr.io/distroless/base

# Copy the built binary and .env file from the builder stage
COPY --from=builder /app/ /

# Set the working directory
WORKDIR /

# Command to run the application
CMD ["/knowledge-manager"]
