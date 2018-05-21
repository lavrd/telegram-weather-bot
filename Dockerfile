FROM golang:1.10.2

WORKDIR /go/src/telegram-weather-bot
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

WORKDIR /go/src/telegram-weather-bot/cmd/bot
RUN go build

CMD ["./bot"]
