/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/16/13
 * Time: 5:35 PM
 */
package main

import "fmt"

func main() {

	var karakter byte = '?'

	switch (karakter){
		case 'A' :  fmt.Println("Karakter A")
		case 'B' :  fmt.Println("Karakter B")
		case '?' :  fmt.Println("Karakter ?")
	}

	var nilai int32 = 91; // nilai silahkan diubah-ubah

	switch {
		case nilai >90 :  fmt.Println("Nilai Anda A")
		case nilai >70 :  fmt.Println("Nilai Anda B")
		case nilai >60 :  fmt.Println("Nilai Anda C")
		default :  fmt.Println("Nilai Anda D")
	}

	for i:=0; i<10;i++{
		switch {
			case i%2 == 0 : fmt.Println("Nilai i % 2 = ",i)
			case i%5 == 0 : fmt.Println("Nilai i % 5 = ",i)
		}
	}

}
