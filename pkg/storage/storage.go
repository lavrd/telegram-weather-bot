package storage

const DBName = "telegram-weather-bot"

type Storage interface {
	Close() error
}
