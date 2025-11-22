# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:latest AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cc-wc .

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/cc-wc .

ENTRYPOINT ["./cc-wc"]
