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
		las, _ := strconv.Atoi(strings.Trim(strings.Trim(s, "\n"), "\r"))
		str,_:=a.ReadString('\n')
		str = strings.Trim(strings.Trim(str, "\n"), "\r")
		mstr := make([]int, las)
		mstra := strings.Split(str," ")
		for j:=0; j<las; j++{
			mstr[j], _ = strconv.Atoi(mstra[j])
		}
		fmt.Printf("%d\n",fn(mstr))
	}
}


func fn(a []int) (res int){
	if len(a) == 1{
		return 1
	}

	m:=make(map[int]int)
	max := 1
	for i:=0; i<len(a); i++{
		_, ok := m[a[i]]
		if ok{
			m[a[i]]++
			if m[a[i]] > max{
				max = m[a[i]]
			}
		}else{
			m[a[i]] = 1
		}
	}
	return max
}