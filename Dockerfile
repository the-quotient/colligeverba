# syntax=docker/dockerfile:1
FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY web/ ./web/
COPY engine/ ./engine/
COPY db/ ./db/

RUN mkdir /data

RUN go build -o /vv web/main.go
EXPOSE 8080
CMD ["/vv"]
