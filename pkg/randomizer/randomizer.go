package randomizer

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandom(values []string) string {
	return values[rand.Intn(len(values))]
}
