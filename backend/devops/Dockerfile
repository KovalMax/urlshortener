FROM golang:1.13.1-alpine3.10 as build

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod go.sum ./

COPY devops/.env .env

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./main"]