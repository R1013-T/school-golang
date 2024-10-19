package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("i is %v, %p\n", i, &i)

	{ // block scope
		i = 12
		fmt.Printf("in block: i is %v, %p\n", i, &i)

		i := 123
		fmt.Printf("in block: i is %v, %p\n", i, &i)

		ii := 2
		fmt.Printf("in block: ii is %v, %p\n", ii, &ii)
	}
	// fmt.Println("ii is", ii) // ii is undefined

	fmt.Printf("i is %v, %p\n", i, &i)
}
