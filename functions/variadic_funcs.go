package main

import "fmt"

func main() {
	result := addEm(0, 0, 0, 10, 10, 10)
	fmt.Println(result)
}

func addEm(nums ...int) (result float32) {
	for _, num := range nums {
		result = result + float32(num)
	}

	return result
}
