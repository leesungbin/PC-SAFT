FROM golang:1.13 as builder
WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server .
RUN ls /app
RUN GOOS=linux go build -a installsuffix -o main .

FROM node:10 as web
WORKDIR /web
COPY client/package.json client/package-lock.json ./
RUN npm install
COPY client/src ./src
COPY client/tsconfig.json ./
RUN npm run build

FROM debian:bullseye-20190708
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=web /web/dist ./web/dist
RUN ls -lah
ENTRYPOINT [ "/root/main" ]