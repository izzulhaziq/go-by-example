package main

import "fmt"

func sumAndMax(values ...int) (int, int) {
	highest := 0
	sum := 0
	for _, value := range values {
		if value > highest {
			highest = value
		}

		sum += value
	}

	return highest, sum
}

func main() {
	values := []int{1, 2, 3, 3, 2, 3, 4, 5, 6, 7, 4, 4, 6, 9, 3}
	max, sum := sumAndMax(values...)

	fmt.Println(max)
	fmt.Println(sum)
}
