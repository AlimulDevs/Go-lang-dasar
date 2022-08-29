package main

import (
	"fmt"
	"strconv"
)

func main() {

	// contoh konversi data 1
	var nim1 int32 = 10000
	var nim2 int64 = int64(nim1)

	var nim3 int16 = int16(nim1)
	var nim4 int8 = int8(nim1)

	fmt.Printf("%T \n", nim1)
	fmt.Printf("%T \n", nim2)
	fmt.Printf("%T \n", nim3)
	fmt.Printf("%T \n", nim4)

	// contoh konversi data 2
	var nama string = "alimul"
	var e byte = nama[2]
	var ambil_e string = string(e)

	fmt.Println(nama)
	fmt.Println(ambil_e)

	// contoh konversi 3 mengubah dari int ke angka

	ress, _ := strconv.Atoi("angka")
	fmt.Printf("%T \n", ress)

}
