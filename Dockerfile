# Use the official Golang image version 1.19
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest


# Copy the Go mod and sum files
COPY go.mod go.sum ./
#
## Download all dependencies
RUN go mod tidy
#
## Copy the source code into the container
#COPY . .
#
## Build the Go application
#RUN go build -o main .

# Expose the desired port (change it if your application listens on a different port)
# EXPOSE 3000

# Set the command to run the binary
# CMD ["./main"]
