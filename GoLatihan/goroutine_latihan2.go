package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func proses(menu string, jlh int) {
	for i := 0; true; i++ {
		a := time.Duration(rand.Intn(5) * 1e10)
		time.Sleep(a)

		fmt.Printf("%v : %v , waktu yang dibutuhkan : %v \n", menu, i, a)
	}
}

func main() {

	for i := 0; i < 100000; i++ {
		go proses("Pesan Jus Jambu", 10)
		go proses("Pesan Nasi Kuning", 10)
		go proses("Pesan Jus Apel", 5)
		go proses("Pesan Bubur Manado", 5)
	}

	fmt.Println(runtime.NumGoroutine())

	var input string
	fmt.Scanln(&input)
	fmt.Println("Selesai!")

}
