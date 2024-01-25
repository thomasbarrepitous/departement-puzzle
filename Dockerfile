# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install templ tool
RUN go install github.com/a-h/templ/cmd/templ@latest

# Generate the Go files from the .templ
RUN templ generate

# Download all the dependencies
RUN go mod tidy
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
