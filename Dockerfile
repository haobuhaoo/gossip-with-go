# Use Go 1.25.5 image
FROM golang:1.25.5

# Set working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Install goose and run migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN goose up

# Build the Go binary
RUN go build -tags netgo -ldflags '-s -w' -o app

# Expose port
EXPOSE 10000

# Start the app
CMD ["./app"]
