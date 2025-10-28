package utils

import "strings"

type Step uint8

const (
	CREATED Step = iota
	UPDATED
)

func IsBlank(data string) bool {
	return len(strings.TrimSpace(data)) == 0
}
