/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/18/13
 * Time: 9:40 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import ("fmt"; "bufio"; "os"; "strings";)

func main() {
	fmt.Printf("Kitik nama Anda : ")
	scn := bufio.NewReader(os.Stdin);
	data, _ := scn.ReadString('\n');
	fmt.Printf("Selamat datang %v!", strings.TrimRight(data,"\n"))
}

