package main

import "fmt"

func WordGenerator(words []string) func() string {
	counter := 0
	return func() string {
		w := words[counter]
		counter += 1
		if counter == len(words) {
			counter = 0
		}
		return w
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
