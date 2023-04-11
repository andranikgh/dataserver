# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files into the container
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Set the command to run the executable
CMD ["./main"]
