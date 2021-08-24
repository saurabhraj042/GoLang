package main

import "fmt"

const englishHelloAgainPrefix = "Hello, "
const frenchHelloAgainPrefix = "Bonjour, "
const spanishHelloAgainPrefix = "Hola, "
const french = "French"
const spanish = "Spanish"

func HelloAgain(name string, language string) string{
	if name == ""{
		name = "World"
	}

	prefix := englishHelloAgainPrefix

	switch language {
	case french:
		prefix = frenchHelloAgainPrefix
	case spanish:
		prefix = spanishHelloAgainPrefix
	}

	return prefix + name
}

func Hello() string{
	return "Hello World from outside the main() scope"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Hello())
}