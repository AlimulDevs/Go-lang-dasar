package main

import "fmt"

func main() {
	var nilai_1 = 65
	var nilai_2 = 90

	var lulus_nilai_1 bool = nilai_1 >= 67
	var lulus_nilai_2 bool = nilai_2 >= 75
	var lulus bool = lulus_nilai_1 && lulus_nilai_2
	var lulus2 bool = lulus_nilai_1 || lulus_nilai_2

	fmt.Println(lulus)
	fmt.Println(lulus2)
}
