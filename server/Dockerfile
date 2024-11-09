# syntax=docker/dockerfile:1

FROM golang:1.23.0

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY  . .

RUN go build -o main .

EXPOSE 3000

CMD ["/main"]
