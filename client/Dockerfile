FROM golang:1.14

WORKDIR /app

COPY . /app
RUN go build

ENTRYPOINT client
