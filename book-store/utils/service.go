package utils

import (
	"math/rand"
	"strings"
)

var letter []string

func CreateLettersList() {
	for ch := 'a'; ch <= 'z'; ch++ {
		letter = append(letter, string(ch))
	}
}

func GenerateID() string {
	var id []string
	for _ = range 4 {
		random := rand.Intn(len(letter))
		id = append(id, letter[random])
	}
	uuid := strings.Join(id, "")
	return uuid
}
