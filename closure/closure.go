package main

import "fmt"

func main() {
	angka := 0

	deklarasi := func() {
		fmt.Println("deklarasi")
		angka++
	}

	deklarasi()
	deklarasi()
	deklarasi()
	fmt.Println(angka)

}
