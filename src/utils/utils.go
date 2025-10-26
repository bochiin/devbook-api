package utils

import "strings"

func IsBlank(data string) bool {
	return len(strings.TrimSpace(data)) == 0
}
