// Package iteration ...
package iteration

// Repeat character 5 times
func Repeat(character string) string {
	var repeated string

	for range 5 {
		repeated += character
	}

	return repeated
}
