package docker

import (
	"github.com/docker/docker/pkg/namesgenerator"
	"strings"
)

func Generate() string {
	nameParts := strings.Split(namesgenerator.GetRandomName(0), "_")

	for i := 0; i < len(nameParts); i++ {
		nameParts[i] = strings.Title(nameParts[i])
	}

	return strings.Join(nameParts, " ")
}

func GenerateMultiple(count int) []string {
	s := make([]string, count)
	for i := 0;  i <= count; i++ {
		s[i] = Generate()
	}

	return s
}
