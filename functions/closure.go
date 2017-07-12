package main

import "fmt"

func sayName(name string) func() {
	return func() {
		fmt.Println("Hello " + name)
	}
}

func highfive(name string) func(int) string {
	count := 0
	limit := 33
	return func(times int) string {
		for i := 0; i < times; i++ {
			if count > limit {
				fmt.Println("You have reched your high five limit!")
			} else {
				fmt.Println("High five " + name + "!")
			}

			count++
		}

		return "yeay"
	}
}

func main() {
	sayAnjana := sayName("Anjana")
	sayMichelle := sayName("Michelle")

	sayAnjana()
	sayMichelle()

	var result string
	highfiveMichelle := highfive("Michelle")
	result = highfiveMichelle(5)
	fmt.Println(result)

	result = highfiveMichelle(10)
	fmt.Println(result)

	highfiveMichelle(20)
	fmt.Println(result)
}
