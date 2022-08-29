package main

import "fmt"

func main() {
	namaAwal, namaTengah, namaAkhir := nama()

	fmt.Println(namaAwal, namaTengah, namaAkhir)

}

func nama() (firstName string, midName string, lastName string) {
	firstName = "Nur"
	midName = "alimul"
	lastName = "haq"

	return
}
