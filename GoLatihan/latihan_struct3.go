package main

import "fmt"

type OrangTua struct{
	nama_ayah string
	usia_ayah int
	nama_ibu string
	usia_ibu int
}

type Siswa struct {
	nama string
	nomor_induk int
	berat float32
	tinggi float32
	ortu OrangTua
}

func (s *Siswa) ToString(){
	fmt.Println("Nama Siswa : ", s.nama)
	fmt.Println("Nomor Induk : ", s.nomor_induk)
	fmt.Println("Berat Badan : ", s.berat)
	fmt.Println("Tinggi Badan :", s.tinggi)
	fmt.Println("Nama Ayah :", s.ortu.nama_ayah)
	fmt.Println("Usia Ayah:", s.ortu.usia_ayah)
	fmt.Println("Nama Ibu :", s.ortu.nama_ibu)
	fmt.Println("Usia Ibu :", s.ortu.usia_ibu)
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

	ayash := new(Siswa)
	ayash.nama ="Abdullah Ayash"
	ayash.ortu.nama_ayah = "Abdul Rizal Adompo"

	// memanggil fungsi ToString untuk menampilkan data
	rama.ToString()
	fmt.Println("\n")
	ayash.ToString()

}
