package model

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrConflict      = errors.New("conflict error")
	ErrPartsNotFound = errors.New("parts not found")
)
