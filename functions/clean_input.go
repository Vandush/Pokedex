package functions

import (
	"strings"
)

func CleanInput(text string) []string {
	var list []string
	list = strings.Fields(text)
	for i, text := range list {
		list[i] = strings.ToLower(text)
	}
	return list
}
