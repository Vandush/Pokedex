package main

import "strings"

func CleanInput(text string) []string {
	var list []string
	list = strings.Fields(text)
	for i, string := range list {
		list[i] = strings.ToLower(string)
	}
	return list
}
