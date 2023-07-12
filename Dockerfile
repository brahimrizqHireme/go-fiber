# Use the official Golang image version 1.19
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app


# Copy the Go mod and sum files
COPY go.mod go.sum ./
#
RUN ls -al .
## Download all dependencies
RUN go mod tidy

RUN go install \
    github.com/cosmtrek/air@latest

# Install gocode-gomod for Go autocompletion (required for Go modules)
RUN GO111MODULE=on go install github.com/stamblerre/gocode@latest \
    && mv /go/bin/gocode /go/bin/gocode-gomod

# Set environment variables for Go module support and enabling Go modules in GOPATH mode
ENV GO111MODULE=on \
    GOFLAGS=-mod=vendor

## Copy the source code into the container
#COPY . .
#
## Build the Go application
#RUN go build -o main .

# Expose the desired port (change it if your application listens on a different port)
# EXPOSE 3000

# Set the command to run the binary
# COPY . .

# # Set the command to run the application
# CMD ["go", "run", "main.go"]