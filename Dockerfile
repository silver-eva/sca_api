# Use the official Golang image as the base
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o api

# Use a minimal base image for the final build
FROM scratch

# # Copy the compiled binary and necessary files from the builder stage
COPY --from=builder /app/api ./api
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./api"]
