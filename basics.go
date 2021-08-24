package main

import "fmt"

const helloAgainPrefix = "Hello, "

func HelloAgain(name string) string{
	if name == ""{
		name = "World"
	}

	return helloAgainPrefix + name
}

func Hello() string{
	return "Hello World from outside the main() scope"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Hello())
}