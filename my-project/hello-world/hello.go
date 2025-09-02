package main

import "fmt"

// try to make constants to capure meaning of values (can aid performance too)
const (
	spanish = "Spanish"
	french  = "French"
	
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// starts uppercase for public functions to be exported
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// starts lowercase for private functions to not be exported
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // no need to add prefix because of our func siganture preset that
}

func main() {
	fmt.Println(Hello("", ""))
}
