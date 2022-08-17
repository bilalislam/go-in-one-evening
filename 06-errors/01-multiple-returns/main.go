package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

/*
 * you need to add another pair of parantheses and seperate the types with commas (,)
 * and you can return the same number of variables and seperate them unchanged but in reverse order
 */
func Swap(x, y string) (string, string) {
	return y, x
}
