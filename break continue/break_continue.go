package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {

		if i == 5 {
			break
		}
		fmt.Println("halaman", i)
	}

	continues()

}

func continues() {
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println("tes", i)
		}

	}

}
