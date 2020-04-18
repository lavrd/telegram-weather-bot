package storage

import (
	"telegram-weather-bot/pkg/types"

	"github.com/pkg/errors"
)

const DBName = "telegram-weather-bot"

var (
	ErrUserNotFound = errors.New("user not found")
)

type Storage interface {
	GetUser(telegramID int64) (types.User, error)
	UpdateUserUnits(telegramID int64, units string) error

	Close() error
}
