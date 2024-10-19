package main

import "fmt"

func main() {
	// 配列
	x := [3]int{1, 2, 3}
	// x.append(4)
	fmt.Printf("x is %v, %T\n", x, x)

	// スライス
	y := []int{1, 2, 3}
	y = append(y, 4)
	fmt.Printf("y is %v, %T\n", y, y)
}
