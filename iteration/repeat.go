// Package iteration ...
package iteration

import "strings"

// Repeat character 5 times
func Repeat(character string, count int) string {
	var repeated strings.Builder

	for range count {
		repeated.WriteString(character)
	}

	return repeated.String()
}
