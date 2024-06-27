# syntax=docker/dockerfile:1

# Stage 1: Build the Go application
FROM golang:1.22.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Run the Go application
FROM gcr.io/distroless/base-debian10

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]