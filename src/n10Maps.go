package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["joker1"] = 7
	m["joker2"] = 8

	fmt.Println("map:", m)

	var k1 = m["joker"]

	fmt.Println("k1:", k1)

	delete(m, "joker")

	_, prs := m["joker2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(n)

}
