package main

import "fmt"

func main() {
	if n := 9; n < 0 {
		fmt.Println("n is negative")
	} else if n > 10 {
		fmt.Println("n is more than 10")
	} else if n > 0 && n < 10 {
		fmt.Println("but actually n is between 0 and 10")
	}
}
