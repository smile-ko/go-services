FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

CMD ["go", "run", "./cmd/server"]

EXPOSE 8002