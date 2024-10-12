package main

import "fmt"

func main() {
	fmt.Println("Hello, Go World 💙\n")

	var i int = 10 // 整数型
	printValueAndType("i", i)

	var ui uint = 20 // 符号なし整数型
	printValueAndType("ui", ui)

}

func printValueAndType(name string, v interface{}) {
	fmt.Printf("%s: %v (%T)\n", name, v, v)
}
