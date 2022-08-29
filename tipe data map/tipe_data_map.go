package main

import "fmt"

func main() {
	// ex 1

	angka := map[string]int{
		"angka1": 1234,
		"angka2": 5678,
	}
	fmt.Println(angka["angka1"])
	fmt.Println(angka["angka2"])

	fmt.Println(angka)

	delete(angka, "angka1")
	fmt.Println(angka)

}
