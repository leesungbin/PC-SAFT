FROM golang:1.13 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . ./
RUN GOOS=linux go build -a -installsuffix cgo -o main github.com/leesungbin/PC-SAFT/
