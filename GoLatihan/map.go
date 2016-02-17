/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 12/1/13
 * Time: 6:19 AM
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
	bank_soal := make(map[string]string)  // map[key]value
	bank_soal["s1"] = "Presiden RI 1 adalah..."
	bank_soal["s1a"] = "Soekarno"
	bank_soal["s1b"] = "Soeharto"
	bank_soal["s1c"] = "BJ Habibie"
	bank_soal["s1d"] = "Gusdur"
	bank_soal["s1e"] = "Sajarwo"
	bank_soal["s1jawaban"] = "s1a"

	delete(bank_soal,"s1e") // menghapus elemen

	fmt.Println("Length Bank Soal :", len(bank_soal))
	fmt.Println(bank_soal)

	for k, v := range bank_soal {
		fmt.Printf("\nKey: %v, Val: %v", k, v)
	}

	if soal1, ok := bank_soal["s1"]; ok {
		fmt.Println("\n\nSoal : ", soal1, "\nStatus : ", ok)
	}

	soal2, ok := bank_soal["s2"]
	fmt.Println("\nSoal : ", soal2, "\nStatus : ", ok)
}

