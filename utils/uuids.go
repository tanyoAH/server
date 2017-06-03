package utils

import "github.com/satori/go.uuid"

func GenerateUUIDv4() string {
	return uuid.NewV4().String()
}
