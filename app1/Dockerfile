
# Use official Golang image as base
FROM golang:1.18-alpine

WORKDIR /app

# Copy the Go application code
COPY . .

# Build the Go application
RUN go mod init app1
RUN go mod tidy
RUN go build -o app1

# Expose the port and run the app
EXPOSE 8001
CMD ["./app1"]
