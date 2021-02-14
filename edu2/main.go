package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	a:=bufio.NewReader(os.Stdin)
	l, _ :=a.ReadString('\n')
	l=strings.Trim(strings.Trim(l,"\n"),"\r")
	ln, _ :=strconv.Atoi(l)
	for i:=0; i<ln;i++{
		s,_:=a.ReadString('\n')
		d, _ :=strconv.Atoi(strings.Trim(strings.Trim(s,"\n"), "\r"))
		fmt.Printf("%s\n",fn(d))

	}

}

func fn(n int) string{
	x:=n/2020
	o:=n%2020
	if o == 0 || x >= o{
		return "YES"
	}
	return "NO"
}