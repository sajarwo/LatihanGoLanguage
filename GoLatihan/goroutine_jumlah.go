package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Jumlah Goroutine : ", runtime.NumGoroutine())
}
