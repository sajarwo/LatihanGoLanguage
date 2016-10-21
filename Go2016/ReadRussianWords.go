package main

import (
	"fmt"
	"os"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/charmap"
	"strings"
)

func main() {
	readfile("C:/Users/SNA/Downloads/Compressed/words.num") // sesuaikan dengan tempat penyimpanan
}

// contoh membaca file dengan encoding tertentu (dalam hal ini bahasa rusia)
// file dapat di download melalui http://www.artint.ru/projects/frqlist/words.num.zip
func readfile(src string){
	f, _ := os.Open(src)
	defer f.Close();
	r := transform.NewReader(f, charmap.Windows1251.NewDecoder())
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		f:=strings.Fields(sc.Text())
		fmt.Println(f[2])
	}
}
