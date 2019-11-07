package helper

import (
	"math/rand"
	"time"
)

// GenerateAnswer return
func GenerateAnswer() string {
	answers := []string{"A", "B", "C", "D"}

	rand.Seed(time.Now().UnixNano())
	return answers[rand.Intn(4)]
}
