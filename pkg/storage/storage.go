package storage

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type Storage interface {
	Upsert(user *User) error
	GetUser(telegramID int64) (*User, error)

	Close() error
}

type User struct {
	TelegramID int64
	Location   string
	Lang       string
	Lat        float64
	Lon        float64
	Units      Units
}

type Units int

const (
	SI Units = iota + 1
	US
)
