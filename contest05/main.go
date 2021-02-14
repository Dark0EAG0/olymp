package main

import (
	"bufio"
	"fmt"
	"math"
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
		las := strings.Split(strings.Trim(strings.Trim(s, "\n"), "\r")," ")
		n, _ :=strconv.ParseFloat(las[0],64)
		k, _ :=strconv.ParseFloat(las[1],64)
		fmt.Printf("%d\n",fn(k,n))
	}
}

func fn(k float64, n float64) int64{
	p:=k
	for p/n < 1{
		p=p+k
	}
	return int64(math.Ceil(p/n))
}