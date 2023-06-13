FROM golang:1.20-alpine

LABEL maintainer="phamhunggl721@gmail.com"

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

CMD ["air", "-c", ".air.toml"]