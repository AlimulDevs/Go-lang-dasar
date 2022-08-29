package main

import "fmt"

func main() {
	runApp(0)
}

func logging() {
	fmt.Println("selesai memanggil")
}

func runApp(value int) {
	defer logging()
	fmt.Println("run app")
	result := 10 / value
	fmt.Println("result", result)
}
