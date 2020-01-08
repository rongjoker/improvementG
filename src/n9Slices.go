package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "1"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)

}
