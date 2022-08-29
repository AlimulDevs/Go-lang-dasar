package main

import "fmt"

func main() {
	// ex 1
	var angka int = 2

	for angka <= 5 {

		fmt.Println("halaman", angka)
		angka++
	}

	// ex 2
	for angka := 3; angka < 7; angka++ {
		fmt.Println("jumlah", angka)
	}

	// ex 3
	nama := []string{
		"nur", "alimul", "haq",
	}

	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])

	}

	// ex 4
	info := make(map[string]string)
	info["nama"] = "nur alimul haq"
	info["nim"] = "200180100"

	for key, value := range info {
		fmt.Println(key, "=", value)
	}
}
