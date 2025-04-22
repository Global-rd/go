# Use the official Golang image as the base image
FROM golang:1.24.2-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY . /app

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o rest-api .

# Expose the port the application runs on
EXPOSE 8000

# Command to run the application
CMD ["./rest-api"]