# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o service .

# Stage 2: final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/service .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# set the entrypoint command
ENTRYPOINT ["./service"]