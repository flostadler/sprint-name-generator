package randomizer

import (
	"math/rand"
)

func GetRandom(values []string) string {
	return values[rand.Intn(len(values))]
}
