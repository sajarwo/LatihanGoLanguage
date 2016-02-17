package main

import "fmt"


type Politeknik struct{
	nama string;
}

func (p *Politeknik)GetNama() string{
	return p.nama
}

type Institut struct{
	nama string;
}

func (p *Institut)GetNama()string{
	return p.nama
}

type PerguruanTinggi interface {
	GetNama() string
}

func main() {
	// penggunaan pada umumnya
	poli := new(Politeknik)
	poli.nama = "Politeknik Elektronika Negeri Surabaya"

	inst := new(Institut)
	inst.nama = "Institut Teknologi Bandung"

	fmt.Println("Daftar Perguruan Tinggi 1")
	fmt.Println("Nama Politeknik : ", poli.GetNama())
	fmt.Println("Nama Institut : ", inst.GetNama())


	// penggunaan interface
	var pt PerguruanTinggi = poli
	var pt1 PerguruanTinggi = inst

	fmt.Println("\nDaftar Perguruan Tinggi 2")
	fmt.Println("Nama PT (Politeknik) : ", pt.GetNama())
	fmt.Println("Nama PT (Institut) : ", pt1.GetNama())

	// penggunaan interface ketika melewatkan object atau tipe data tertentu
	fmt.Println("\n");
	ToString(poli)
	ToString(inst)
}

func ToString(pt PerguruanTinggi){
	fmt.Println("Nama Pergruan Tinggi : ", pt.GetNama())
}
