# Start from a small, secure base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
#COPY go.mod go.sum ./

# Download the Go module dependencies
#RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./main.go

# Create a minimal production image
FROM alpine:latest

# It's essential to regularly update the packages within the image to include security patches
RUN apk update && apk upgrade \
&& apk add --update sudo \
&& apk add --update curl \
&& apk add --update tcpdump \
&& apk add --update bind-tools \
&& apk add --update bash

# Reduce image size
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
# RUN adduser -D appuser
# USER appuser

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .

# Set any environment variables required by the application
# ENV HTTP_ADDR=:8080

# Expose the port that the application listens on
# EXPOSE 80

# Run the binary when the container starts
CMD ["./app"]