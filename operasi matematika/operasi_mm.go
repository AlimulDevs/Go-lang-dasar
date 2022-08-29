package main

import "fmt"

func main() {
	var a = 1
	var b = 2
	var c = a + b
	var d = c - a
	var e = b * b
	var a2 float32 = 2
	var c2 float32 = 3
	var f float32 = c2 / a2
	var persen = 3 % 6
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(persen)

	// contoh 2

	var i = 10
	i += 10

	fmt.Println(i)
	fmt.Println(+i)

}
