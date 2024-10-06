# syntax=docker/dockerfile:1

FROM golang:1.23.0

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY  . .

#generate go struct for configured databases
RUN sqlc generate

RUN go build -o main .

EXPOSE 8080

CMD ["/main"]
