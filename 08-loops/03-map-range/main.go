package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Values(products map[int]string) []string {
	var values []string
	for _, value := range products {
		values = append(values, value)
	}
	return values
}

func Keys(products map[int]string) []int {
	var keys []int
	for key, _ := range products {
		keys = append(keys, key)
	}

	return keys
}
