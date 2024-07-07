# Use the latest version of Golang as a parent image dynamically
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum to ensure dependencies are downloaded efficiently
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire source code from the current directory to the working directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the builder stage to the /app directory in the production image
COPY --from=builder /app/main .

# Expose port 50051 to the outside world (adjust if your application listens on a different port)
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
