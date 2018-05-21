FROM golang:alpine

WORKDIR /go/src/telegram-weather-bot
COPY . .

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

CMD ["bot"]
