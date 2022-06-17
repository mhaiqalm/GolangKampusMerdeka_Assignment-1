package main

import (
	"fmt"
)

func biodata(nama, alamat, pekerjaan string) {
	fmt.Println("Nama : ", nama)
	fmt.Println("Alamat : ", alamat)
	fmt.Println("Pekerjaan : ", pekerjaan)
	return
}

func main() {
	nama := "Haiqal"
	alamat := "Garut"
	pekerjaan := "Mahasiswa"
	biodata(nama, alamat, pekerjaan)
}

type biodata1 struct {
	nama1      string
	alamat1    string
	pekerjaan1 string
}

func main1() {
	var info biodata1
	info.nama1 = "Alex"
	info.alamat1 = "St Petersburg"
	info.pekerjaan1 = "Programmer"

	fmt.Println("Nama : ", info.nama1)
	fmt.Println("Alamat : ", info.alamat1)
	fmt.Println("Pekerjaan : ", info.pekerjaan1)