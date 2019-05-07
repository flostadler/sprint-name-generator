package randomizer

import (
	"github.com/docker/docker/pkg/random"
)

func GetRandom(values []string) string {
	return values[random.Rand.Intn(len(values))]
}
