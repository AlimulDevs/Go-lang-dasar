package main

import "fmt"

func main() {
	var nama Mahasiswa
	nama.nama = "nur alimul haq"
	nama.nim = 200180100
	nama.alamat = "laut tador"

	fmt.Println("nama :", nama.nama)
	fmt.Println("nim :", nama.nim)
	fmt.Println("alamat :", nama.alamat)
	hallo(nama, "loool")

	type Pegawai struct {
		nama  string
		usia  int
		agama string
	}

	DataDiri := Pegawai{
		nama:  "andi",
		usia:  12,
		agama: "islam",
	}

	fmt.Println(DataDiri.nama)
	fmt.Println(DataDiri.usia)
	fmt.Println(DataDiri.agama)

}

type Mahasiswa struct {
	nama   string
	nim    int
	alamat string
}

func hallo(mahasiswa Mahasiswa, isian string) {
	fmt.Println(isian, mahasiswa.nim)

}
