FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8082