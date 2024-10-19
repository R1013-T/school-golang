package main

import "fmt"

func main() {
	var i int = 10 // 整数型
	print("i", i)

	var ui uint = 20 // 符号なし整数型
	print("ui", ui)

	ii := 123 // 省略できる(型推論)
	print("ii", ii)

	ff := 3.14 // 浮動小数点型
	print("ff", ff)
}

func print(name string, v interface{}) {
	fmt.Printf("%s: %v (%T)\n", name, v, v)
}
