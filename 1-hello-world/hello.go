package main

import "fmt"

const (
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "

	french            = "French"
	frenchHelloPrefix = "Bonjour, "

	englishHelloPrefix = "Hello, "
)

// string keyword indicates the func returns string
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// prefix variable type string gets returned
// switch variable
// case value inside variable
// default when none of the case matches
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Elv", "English"))
}
