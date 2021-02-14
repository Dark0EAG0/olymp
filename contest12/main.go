package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := bufio.NewReader(os.Stdin)
	l, _ := a.ReadString('\n')
	l = strings.Trim(strings.Trim(l, "\n"), "\r")
	ln, _ := strconv.Atoi(l)
	for i := 0; i < ln; i++ {
		s, _ := a.ReadString('\n')
		las := strings.Split(strings.Trim(strings.Trim(s, "\n"), "\r"), " ")
		n, _ := strconv.ParseInt(las[0], 10, 64)
		k, _ := strconv.ParseInt(las[1], 10, 64)

		fmt.Printf("%d\n", fn(int64(n), int64(k)))
	}
}

func fn(a int64, b int64) (C int64) {
	if b == 1 {
		C++
		b++
	}
	if a < b{
		return 1+C
	}
	bcount:=int64(1)
	c:=howlong(a, b)
	for bcount < 10{
		c=min(howlong(a, b+bcount) + bcount, c)
		bcount ++
	}
	return C+c
}

func howlong(a int64, b int64) (r int64){
	for a > 0{
		a/=b
		r++
	}
	return r
}

func min(a int64, b int64) int64{
	if a < b{
		return a
	}
	return b
}