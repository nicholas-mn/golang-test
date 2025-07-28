package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	danish             = "Danish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	danishHelloPrefix  = "Hej, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case danish:
		prefix = danishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Nicholas", ""))
}
