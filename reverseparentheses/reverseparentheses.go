// Package reverseparentheses ...
package main

import (
	"strings"
)

func Reverse(origin string) (reversedString string) {
	reversedString = origin
	if !hasParantheses(reversedString) {
		return reversedString
	}

	for hasParantheses(reversedString) {
		reversedString = splitReverseAndRemoveOuterParantheses(reversedString)
	}
	return reversedString
}

func splitReverseAndRemoveOuterParantheses(str string) string {
	startCut := strings.LastIndex(str, "(")
	// find end endCut
	cutStr := str[startCut:]
	endCut := strings.Index(cutStr, ")") + startCut

	toBeReversed := ""
	for i := startCut; i <= endCut; i++ {
		toBeReversed += string(str[i])
	}

	oldStr := toBeReversed

	toBeReversed = reverseString(toBeReversed)

	// remove outer reversed parentheses
	splitStr := strings.Split(toBeReversed, "")
	splitStr = splitStr[1 : len(splitStr)-1]
	toString := strings.Join(splitStr, "")

	outputStr := strings.ReplaceAll(str, oldStr, toString)

	return outputStr
}

func hasParantheses(str string) bool {
	if strings.Index(str, "(") != -1 && strings.LastIndex(str, ")") != -1 {
		return true
	}
	return false
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
