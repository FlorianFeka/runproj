package utils

import "strings"

// Find a string in an array of string
func Find(arr []string, str string) (string, bool) {
	for _, a := range arr {
		if strings.EqualFold(a, str) {
			return a, true
		}
	}
	return "", false
}