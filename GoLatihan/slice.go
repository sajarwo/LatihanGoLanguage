/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 12/1/13
 * Time: 5:04 AM
 * To change this template use File | Settings | File Templates.
 */
package main

import "fmt"

func main() {
	koordinat_x := make([]int32, 5)
	koordinat_y := []int32{ 1, 2, 3, 4, 5, 6 }

	for i := 0; i < 5; i++ {
		koordinat_x[i] = int32(i)
	}

	sy1 := koordinat_y[:]    // [1 2 3 4 5 6]
	sy2 := koordinat_y[2:5]  // [3 4 5]
	sy3 := koordinat_y[3:]   // [4 5 6]
	sy4 := koordinat_y[:3]   // [1 2 3]

	sx_append := append(koordinat_x, 11, 15, 20)  // menambah elemen

	sy_copy := make([]int32, 5)
	copy(sy_copy, koordinat_y[2:5])  // [3 4 5]

	//hilangkan komentar berukut dan amati apa yang terjadi
	// koordinat_y[4]=40;

	fmt.Println("koordinat_x", koordinat_x)
	fmt.Println("koordinat_y", koordinat_y)
	fmt.Println("sy1 ", sy1)
	fmt.Println("sy2 ", sy2)
	fmt.Println("sy3 ", sy3)
	fmt.Println("sy4 ", sy4)
	fmt.Println("sx_append", sx_append)
	fmt.Println("sy_copy", sy_copy)
}
