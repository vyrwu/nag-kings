# BUILD APP
FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY swaggerui/ ./swaggerui/
COPY cert/ ./cert/
COPY go/ ./go/
COPY main.go .

CMD go test ./go/

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

EXPOSE 8081

CMD ["./main"]

# TODO: FROM scratch