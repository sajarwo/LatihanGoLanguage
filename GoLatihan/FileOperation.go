/**
 * Created with IntelliJ IDEA.
 * User: SNA8
 * Date: 10/23/13
 * Time: 10:00 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"os"
	"fmt"
	"strings"
)
//import "strings"

func main() {
	var scontainer string
	buf := make([] byte , 1024)

	f, _ := os.Open("d:/go/file/doc.txt");
	defer f.Close()

	for {
		n, _ := f.Read(buf);
		if  n == 0 { break }
		scontainer= scontainer + string(buf[:n])
		os.Stdout.Write(buf[:n])
	}

	/** or using this for read just one time (didn't define byte-size)
	stat, _ := f.Stat()
	buff := make([]byte, stat.Size())
	f.Read(buff)
	scontainer = string(buff)

	//or we can shorter one
	scontainer = string(f)
	*/

	scontainer = clearSomeCharacter(scontainer,",.;:()[]\n");
 	scontainer1 := splitData(scontainer)
	countWords(scontainer1)
}

// clearing some character
func clearSomeCharacter(dt string, clearchar string) (string) {
	for _,val := range clearchar{
		dt= strings.Replace(dt,string(val),"",-1)
	}
	return dt
}

// split data into array
func splitData(dt string)[]string{
	p1 := strings.Split(dt, " ");

	fmt.Printf("%v",dt);

	for _, val := range p1{
		fmt.Printf("%v \n", val);
	}

	fmt.Printf("Total words : %v \n\n\n", len(p1));
	return p1
}

// count words appearing
func countWords(dt []string)(map[string]int){
	cword := map[string]int{};

	for _, val := range dt{
		if strings.TrimSpace(val) !=""{
			_, check := cword[string(val)]
			if check {
				cword[string(val)] +=1;
			}else{
				cword[string(val)] = 1
			}
		}
	}

	for i, val := range cword{
		fmt.Printf("%v : %v \n", i,  val);
	}
	fmt.Printf("\nThis Map Contains %v words.\n", len(cword));
	return cword
}


