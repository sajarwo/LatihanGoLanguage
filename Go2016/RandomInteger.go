package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  // random nilai integer mulai dari 0-100 
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Print(r.Intn(100))
}
