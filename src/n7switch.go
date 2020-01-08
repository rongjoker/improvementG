package main

import (
	"fmt"
)

func main() {

	i := 2
	fmt.Printf("write", i, "as")

	switch i {
	case 1:
		fmt.Printf("i am one")
	case 2:
		fmt.Println("i am two")
	case 3:
		fmt.Println("i am three")

	default:
		fmt.Println("i am banana")

	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am nool")

		case int:
			fmt.Println("i am int")

		default:
			fmt.Println("i have no idea:", t)

		}
	}

	whatIam(true)
	whatIam(1)
	whatIam("hey")

}
