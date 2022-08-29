package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("app selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("app eror")
	}
	fmt.Println("app berjalan")
}
