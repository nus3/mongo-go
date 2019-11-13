package helper

import (
	"math/rand"
	"time"
)

// GenerateAnswer return
func GenerateAnswer() string {
	answers := []string{"A", "B", "C", "D", "E"}

	rand.Seed(time.Now().UnixNano())
	return answers[rand.Intn(5)]
}
