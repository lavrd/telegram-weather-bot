package storage

import (
	"telegram-weather-bot/pkg/types"

	"github.com/pkg/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Storage interface {
	CreateUser(telegramID int64, lang string) error
	GetUser(telegramID int64) (types.User, error)
	UpdateUserUnits(telegramID int64, units string) error
	UpdateUserLang(telegramID int64, lang string) error

	Close() error
}
