package storage

import (
	"github.com/pkg/errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	TelegramID int64   `gorethink:"telegramID"`
	Location   string  `gorethink:"location"`
	Lang       string  `gorethink:"lang"`
	Lat        float64 `gorethink:"lat"`
	Lon        float64 `gorethink:"lon"`
	Units      string  `gorethink:"units"`
}

type Storage interface {
	GetUser(telegramID int64) (*User, error)
	UpdateUserLang(telegramID int64, lang string) error

	// UpdateUserUnits(telegramID int64, units string) error
	// UpdateUserLang(telegramID int64, lang string) error

	Close() error
}
