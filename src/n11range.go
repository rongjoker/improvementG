package main

import "fmt"

func sumx(num ...int) {
	fmt.Println(num, " ")
	total := 0
	for _, num := range num {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sumx(1, 2, 3)

}
