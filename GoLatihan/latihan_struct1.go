package main

import "fmt"


type Siswa struct {
	nama string
	nomor_induk int
	berat float32
	tinggi float32
}

func main() {
	rama := new(Siswa)
	rama.nama ="Rama Zetta"
	rama.nomor_induk = 701
	rama.berat = 50.2
	rama.tinggi = 165


	fmt.Println("Nama Siswa : ", rama.nama)
	fmt.Println("Nomor Induk : ", rama.nomor_induk)
	fmt.Println("Berat Badan : ", rama.berat)
	fmt.Println("Tinggi Badan :", rama.tinggi)

}
