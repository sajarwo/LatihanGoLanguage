/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/15/13
 * Time: 11:48 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
	for i :=0; i<5; i++ {
		fmt.Println("Nilai i = ",i)
	}

	var j int = 0;
	for  j<5 {
		fmt.Println("Nilai j = ",j)
		j++
	}

	for pos, char := range "САШАРВО" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	for {
		if j++; j <= (10) {
			fmt.Println("Nilai j = ",j)
		}else{
			break
		}
	}
}
