package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	a := [6]int{1, 2, 3, 4, 5, 6}
	for j, v := range a {
		fmt.Printf("%d : %d\n", j, v)
	}

	fmt.Println(a)
}
