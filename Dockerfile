# Start from the official Go image
FROM golang:1.23.5-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create an app directory
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port
EXPOSE 8070

# Command to run the executable
CMD ["./main"]
