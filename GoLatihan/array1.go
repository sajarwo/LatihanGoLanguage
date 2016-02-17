/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 12/1/13
 * Time: 4:23 AM
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
	// contoh 1
	var no_urut [5]int
	no_urut[0] = 701
	no_urut[1] = 702
	no_urut[2] = 703
	no_urut[3] = 704
	no_urut[4] = 705
	fmt.Println("Elemen Array 0: ", no_urut[0])

	// contoh 2
	no_urut1 := [5]int{701, 702, 703, 704, 705}

	fmt.Println("Elemen Array 1: ", no_urut1[1])

	// contoh 3
	var total int32;
	score := [...]int32{10, 20, 15, 3, 18}
	for _, ni := range score {
		total += ni
	}
	fmt.Println("Total Score: ", total)

}
