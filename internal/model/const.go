package model

import (
	"errors"
)

var (
	ErrResourceNotFound = errors.New("resource not found.")
)

const (
	StatusCreated string = "created"
	StatusDeleted string = "deleted"

	ALPHABET     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUMERALS     = "1234567890"
	ALPHANUMERIC = ALPHABET + NUMERALS
)