FROM golang:1.13 as builder
WORKDIR /app
COPY . .
RUN go build -o main .
CMD [ "./saft local serve" ]