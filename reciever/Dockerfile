FROM golang:1.16-alpine AS build

RUN apk add --no-cache git

WORKDIR /src

RUN go mod init rabbitmq-hello-world

RUN go get github.com/streadway/amqp

COPY consumer.go /src

RUN go build consumer.go

FROM alpine as runtime

COPY --from=build /src/consumer /app/consumer

CMD ["/app/consumer"]
