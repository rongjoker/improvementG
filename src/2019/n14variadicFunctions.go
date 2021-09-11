package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
}

func main() {

	sum(1, 2, 3)
	sum(1, 2, 3, 4)

}
