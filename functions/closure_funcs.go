package main

import "fmt"

func main() {
	add := Operate(10, 10, 10)

	add()
	add()
	fmt.Println(add())
}

func Operate(nums ...int) func() float32 {
	result := 0
	times_done := 0
	return func() float32 {
		for _, num := range nums {
			result += num
		}

		times_done += 1
		fmt.Printf("Loaded %v times \n", times_done)
		return float32(result)
	}
}
