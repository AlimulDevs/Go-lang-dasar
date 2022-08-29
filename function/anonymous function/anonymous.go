package main

import "fmt"

func main() {

	namaku := func(nama string) bool {
		return nama == "alimul"

	}

	ortu("misnawati", namaku)
	ortu("lol", func(nama string) bool {
		return nama == "lololol"
	})

}

type Namaku func(string) bool

func ortu(namaOrtu string, namaku Namaku) {
	if namaku(namaOrtu) {
		fmt.Println("ortuku")
	} else {
		fmt.Println("bukan ortuku")
	}
}
