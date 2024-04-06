# Start from a golang image
FROM docker.arvancloud.ir/golang:1.22.1

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.go .

# Expose the port on which the application will run
EXPOSE 8080

# Start the application
CMD ["./main"]
