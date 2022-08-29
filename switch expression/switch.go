package main

import "fmt"

func main() {
	// ex 1
	nama := "alimul"

	switch nama {
	// jika alimul
	case "alimul":
		fmt.Println("hallo alimul")
		// jika bukan alimul
	default:
		fmt.Println("bukan alimul")
	}

	// ex 2
	nim := "200180100"

	switch length := len(nim); length < 7 {
	case true:
		fmt.Println("benar")

	case false:
		fmt.Println("kelebihan angka")

	}

	// ex 3
	kelas := "A5"

	switch {
	case kelas == "A3":
		fmt.Println("kelas benar")

	default:
		fmt.Println("salah kelas")

	}
}
