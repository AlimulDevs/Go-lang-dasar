package main

import "fmt"

func main() {
	namaku := nama("barol")

	fmt.Println(namaku)
}

func nama(namaDepan string) string {
	if namaDepan == "alimul" {
		return "hai alimul"
	} else {
		return "siapa kau"
	}

}
