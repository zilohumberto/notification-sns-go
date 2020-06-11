# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN GO111MODULE=on go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN chmod +x /app

# Build the Go app
RUN cd /app/cmd/notification-server/ && go build -o main .

# Command to run the executable
CMD ["./cmd/notification-server/main"]
