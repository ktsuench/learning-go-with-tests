package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
