package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func proses(nama string, jlh int, c chan string){

	for i:=0;i<jlh;i++{
		a :=time.Duration(rand.Intn(5)*1e9)
		time.Sleep(a)

		// proses mengirimkan data ke main program
		c <- nama+" : "+strconv.Itoa(i)
	}
}

func main() {
	c :=make(chan string)
	go proses("Pesan Soto", 10, c)

	for i:=0;i<10;i++{
		// menunggu data yang dikirimkan melalu channel
		fmt.Println(<-c)
	}
}
