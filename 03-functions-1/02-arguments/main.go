package main

import "fmt"

func main() {
	Greet("Alice")
	Greet("Bob")
}

func Greet(s string) {
	fmt.Printf("Hello, %s", s)
	fmt.Println()
}
