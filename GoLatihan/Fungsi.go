/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/16/13
 * Time: 7:08 PM
 */
package main

import "fmt"

func main() {
	var nilai, err int

	nilai, err = penjumlahan(2, 3)        // silahkan diganti dengan bilangan negatif
	fmt.Println("Hasil = ", nilai)
	fmt.Println("Pesan = ", err)
}

func penjumlahan(a int, b int) (int, int) {
	var err int;
	if a < 0 || b < 0 { // periksa bilangan negatif
		err = -1
	} else {
		err = 0
	}
	return a + b, err;
}
