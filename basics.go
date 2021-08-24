package main

import "fmt"

func Hello() string{
	return "Hello World from outside the main() scope"
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(Hello())
}