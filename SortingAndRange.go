package main

import (
	"fmt"
	"sort"
)

func main() {

	data := []int{1, 3, 5, 4, 7, 6, 9, 8}
	fmt.Println(data)                //[1 3 5 4 7 6 9 8]
	sort.Ints(data)
	fmt.Println(data)                //[1 3 4 5 6 7 8 9]
	fmt.Println("---------------")
	for v := range data {
		fmt.Println(v)                // print index 0-7
	}
	fmt.Println("---------------")
	for k, v := range data {
		fmt.Println(k, v)        // print k,v
	}
}