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
	ns,_:=strconv.Atoi(sc.Text())

	for i:=0; i < ns; i++{
		sc.Scan()
		d, _ :=strconv.Atoi(strings.Split(sc.Text(), " ")[1])
		l, _ :=strconv.Atoi(strings.Split(sc.Text(), " ")[0])
		a := make([]int, l)
		sc.Scan()
		for j,el := range strings.Split(sc.Text()," "){
			a[j],_ = strconv.Atoi(el)
		}
		fmt.Printf("%s\n",fn(a, d))
	}
}

func fn(a []int, d int) string{
	min1 := a[0]
	min2 := a[1]
	max:= a[0]
	if a[0] < a[1]{
		max= a[1]
	}
	for i:=2; i < len(a); i++{
		if a[i] > max{
			max=a[i]
		}
		if a[i] < min1{
			if min1 < min2{
				min2 = min1
			}
			min1 = a[i]
		} else {
			if a[i] < min2 {
				min2 = a[i]
			}
		}
	}
	if max <= d || (min1+min2) <= d{return "YES"}
	return "NO"
}
