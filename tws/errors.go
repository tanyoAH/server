package tws

import "errors"

var (
	ErrInvalidJSON        = errors.New("Error invalid message JSON payload.")
	ErrInvalidId          = errors.New("Error invalid message ID.")
	ErrInvalidMessageType = errors.New("Error invalid message type.")
	ErrMethodNotFound     = errors.New("Error request method not found.")
)
