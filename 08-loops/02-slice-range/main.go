package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(numbers ...int) int {
	var total int
	for _, e := range numbers {
		total += e
	}
	return total
}
