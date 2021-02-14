package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc:=bufio.NewScanner(os.Stdin)
	sc.Scan()
	l,_:=strconv.Atoi(sc.Text())
	for i:=0; i<l;i++{
		sc.Scan()
		a := sc.Text()
		sc.Scan()
		fmt.Printf("%s\n",fn(a, sc.Text()))
	}
}

func fn(a string, d string) string{
	l1:=len(a)
	l2:=len(d)
	p:=l1*l2
	i1,i2 := 1,1
	for true{
		if i1*l1 == i2*l2 {
			if strings.Repeat(a, i1) == strings.Repeat(d, i2)  {
				return strings.Repeat(a, i1)
			}else{
				return "-1"
			}

		}
		if i1*l1>p || i2*l2>p{
			break
		}
		if l2*i2>l1*i1{
			i1++
		}else {
			i2++
		}

	}
	return "-1"
}
