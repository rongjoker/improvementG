package main

import "fmt"

type people struct {
	name string
	age  int
}

func newPeople(name string) *people {

	p := people{age: 10, name: name}
	p.age = 15
	return &p
}

func main() {

	fmt.Println(people{age: 111})
	fmt.Println(people{"233", 16})

	fmt.Println(newPeople("kkkk"))

}
