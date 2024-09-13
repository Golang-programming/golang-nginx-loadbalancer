
# Use official Golang image as base
FROM golang:1.18-alpine

WORKDIR /app

# Copy the Go application code
COPY . .

# Build the Go application
RUN go mod init app2
RUN go mod tidy
RUN go build -o app2

# Expose the port and run the app
EXPOSE 8002
CMD ["./app2"]
