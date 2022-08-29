package main

import "fmt"

func main() {

	// ex 1
	var nama [3]string

	nama[0] = "Nur"
	nama[1] = "Alimul"
	nama[2] = "Haq"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	// ex 2

	var value = [4]int{
		1, 2, 3, 4,
	}

	fmt.Println(value)

	// untuk mengetahui panjang array
	fmt.Println(len(value))
}
