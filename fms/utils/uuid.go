package utils

import (
	"math/rand"
	"strings"
)

func GenerateID() string {
	var charset []string
	for char := 'A'; char <= 'Z'; char++ {
		charset = append(charset, string(char))
	}

	var uuid []string
	for _ = range 5 {
		random := rand.Intn(len(charset))
		uuid = append(uuid, charset[random])
	}

	return strings.Join(uuid, "")
}
