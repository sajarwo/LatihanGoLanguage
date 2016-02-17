/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/15/13
 * Time: 11:24 PM
 */
package main

import "fmt"
import "os"

func main() {
	var x int = 10;
	var y int;


	if x < 20 {
		y = 20;
		fmt.Println("Nilai x =",x)
		fmt.Println("Nilai y =",y)
	}

	//Penggunaan if dengan statement
	if x = 50; x > y {
		fmt.Println("x > y")
	}

	os_name, err := os.Hostname()
	if err == nil {
		fmt.Println("\nOs Name : "+os_name)
	}
}
