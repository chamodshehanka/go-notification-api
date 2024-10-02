FROM golang:1.23.2-alpine3.20 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .


# Final stage
FROM alpine:3.20

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./main"]