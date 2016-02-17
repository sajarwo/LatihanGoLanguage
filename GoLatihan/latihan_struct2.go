package main

import "fmt"


type Siswa struct {
	nama string
	nomor_induk int
	berat float32
	tinggi float32
	ortu OrangTua
}

type OrangTua struct{
	nama_ayah string
	usia_ayah int
	nama_ibu string
	usia_ibu int
}

func main() {
	rama := new(Siswa)
	rama.nama ="Rama Zetta"
	rama.nomor_induk = 701
	rama.berat = 50.2
	rama.tinggi = 165
	rama.ortu.nama_ayah = "Gatot H"
	rama.ortu.usia_ayah = 65
	rama.ortu.nama_ibu = "Eni S"
	rama.ortu.usia_ibu = 55

	fmt.Println("Nama Siswa : ", rama.nama)
	fmt.Println("Nomor Induk : ", rama.nomor_induk)
	fmt.Println("Berat Badan : ", rama.berat)
	fmt.Println("Tinggi Badan :", rama.tinggi)
	fmt.Println("Nama Ayah :", rama.ortu.nama_ayah)
	fmt.Println("Usia Ayah:", rama.ortu.usia_ayah)
	fmt.Println("Nama Ibu :", rama.ortu.nama_ibu)
	fmt.Println("Usia Ibu :", rama.ortu.usia_ibu)

}
