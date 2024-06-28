# Start with the official Go image
FROM golang:1.22.3

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY ../go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code
COPY .. .

# Build the Go application
RUN go build -o /se-school-case

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["/se-school-case"]
