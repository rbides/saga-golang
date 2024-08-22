# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23

COPY . /app
WORKDIR /app
RUN go mod tidy



