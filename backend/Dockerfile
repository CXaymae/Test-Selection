# Use Golang as the base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]
