package model

import "errors"

var (
	ErrFailedInsert  = errors.New("failed to insert")
	ErrFailedUpdate  = errors.New("failed to update")
	ErrOrderNotFound = errors.New("order not found")
)
