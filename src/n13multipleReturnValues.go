package main

import "fmt"

func vals() (int, int) {
	return 3, 8
}

func main() {

	a, b := vals()

	println("a:", a)
	println("b:", b)

	_, c := vals()
	fmt.Println(c)

}
