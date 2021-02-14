package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a:=bufio.NewReader(os.Stdin)
	l, _ :=a.ReadString('\n')
	l=strings.Trim(strings.Trim(l,"\n"),"\r")
	ln, _ :=strconv.Atoi(l)
	for i:=0; i<ln;i++{
		s,_:=a.ReadString('\n')
		d, _ :=strconv.ParseInt(strings.Trim(strings.Trim(s,"\n"), "\r"),10 , 64)
		fmt.Printf("%s\n",fn(d))

	}
}

func fn(n int64) string {
	if n & (n - 1) == 0{
		return "YES"
	}
	return "NO"
}
