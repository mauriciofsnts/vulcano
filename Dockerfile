# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod .
COPY go.sum .
COPY Makefile .

COPY internal internal
COPY cmd cmd 

RUN go mod download

RUN make dist 

EXPOSE 8080

ENTRYPOINT [ "/app/vulcano" ]