FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/discount-service
COPY . .

# Using go get.
RUN go get

# building apps in discount-service
RUN go build -o discount-service

# running discount-service
ENTRYPOINT ./discount-service

# running in port
EXPOSE 8002