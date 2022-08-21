package main

import "fmt"

var counter int

func AllocateBuffer() *string {
	if counter < 3 {
		counter++
		return new(string) // &string so *pointer is equal 0
	}
	return nil
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
