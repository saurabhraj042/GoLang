package main

import "fmt"

const englishHelloAgainPrefix = "Hello, "
const frenchHelloAgainPrefix = "Bonjour, "
const french = "French"

func HelloAgain(name string, language string) string{
	if name == ""{
		name = "World"
	}

	if language == french{
		return frenchHelloAgainPrefix + name
	}

	return englishHelloAgainPrefix + name
}

func Hello() string{
	return "Hello World from outside the main() scope"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Hello())
}