package main

import "fmt"

func main() {
	// ex 1

	var angka = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	var slice1 = angka[0:8]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// ex 2
	newslice := make([]string, 3, 3)
	newslice[0] = "alimul"
	newslice[1] = "haq"
	newslice[2] = "haq"
	fmt.Println(newslice, len(newslice), cap(newslice))

	var foods []string = []string{
		"burger", "kebab", "nasi goreng",
	}

	var target []string = make([]string, 3)

	copy(target, foods)

	fmt.Println(target)

}
