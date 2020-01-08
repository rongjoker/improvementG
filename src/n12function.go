package main

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(111, 2)
	println("res:", res)

	res = plusPlus(1, 2, 3)

	println("res:", res)

}
