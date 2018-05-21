FROM golang:alpine

WORKDIR /go/src/github.com/spacelavr/telegram-weather-bot
COPY . .

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["bot"]
