# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any dependencies required by the application
RUN go get -d -v ./...
RUN go install -v ./...

# Specify the command to run when the container starts
CMD ["go", "run", "main.go"]
