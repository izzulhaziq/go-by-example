package main

import (
	"fmt"
)

func main() {
	s := make([]string, 2)
	s[0] = "solve"
	s[1] = "for"

	e := []string{"happy"}
	s = append(s, e...)

	fmt.Println("s: ", s)

	k := s[:2]
	fmt.Println("k: ", k)
}
