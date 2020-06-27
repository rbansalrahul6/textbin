package models

import (
	"errors"
	"time"
)

// DOUBT: Why not connst?
var (
	ErrNoRecord = errors.New("models: no matching record found")
	// error for invalid credential
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	//error for duplicate email
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	email          string
	HashedPassword []byte
	Created        time.Time
}
