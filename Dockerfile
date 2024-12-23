# Start from the official Go image with specific version
FROM golang:1.23.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code file
COPY main.go .

# Build the Go application
RUN go build -o webserver main.go

# Expose port 80
EXPOSE 80

# Run the compiled binary
CMD ["./webserver"]
