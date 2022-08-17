package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(i1, i2 int) (float64, error) {
	if i2 == 0 {
		return 0, errors.New("divide by zero exception")
	}

	return float64(i1 / i2), nil
}
