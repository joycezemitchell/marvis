# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the working directory
COPY . .

WORKDIR /app/backend

ENV GOPROXY=direct

# Expose the port the server listens on
EXPOSE 9090

# Build the Go application
# RUN go build -o main2 ./cmd/aichat


# Set the entry point for the container
ENTRYPOINT ["./main"]