# Get Go base image
FROM golang:alpine
MAINTAINER Wendell Adriel <wendelladriel.ti@gmail.com>

# Create project structure
COPY . /go/src/github.com/WendellAdriel/go-simple-cli-boilerplate

# Enter the project folder
WORKDIR /go/src/github.com/WendellAdriel/go-simple-cli-boilerplate

# Install necessary tools
RUN apk add --no-cache git mercurial

# Install dependencies
RUN go get ./...

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app
RUN cp app /app

# Remove tools that are not needed anymore
RUN apk del git mercurial

# Set the binary path as the entrypoint
ENTRYPOINT ["/app"]
