package twsproto

import (
	"math/rand"
	"strings"
)

var (
	// TODO assign reserved ids
	reservedIds = []string{}
)

func GenerateRandomId() string {
	return generateRandomStringToken(16)
}

func IsValidId(id string) bool {
	return id != "" && len(id) < 60
}

func IsValidName(name string) bool {
	return len(name) < 255
}

func IsReservedId(id string) bool {
	for _, val := range reservedIds {
		if id == val {
			return true
		}
	}
	return false
}

func IsStdId(id string) bool {
	return strings.HasPrefix(id, "std_")
}

func generateRandomStringToken(lenBytes int) string {
	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, lenBytes)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
