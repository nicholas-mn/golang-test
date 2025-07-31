// Package floyd ...
package floyd

import "strconv"

func Floyd(number int) string {
	triangle := ""

	currentInt := 1
	for i := range number {
		currentLine := ""

		for index := range i + 1 {
			currentLine += strconv.Itoa(currentInt)

			if index != i {
				currentLine += " "
			}

			currentInt++
		}

		triangle += currentLine
		if i+1 != number {
			triangle += "\n"
		}
	}

	return triangle
}
