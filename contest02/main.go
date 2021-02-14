package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd:=bufio.NewReader(os.Stdin)
	tx, _:=rd.ReadBytes('\n')
	l,_ := strconv.Atoi(strings.Trim(strings.Trim(string(tx), "\n"), "\r"))
	for i:=0; i<l;i++{
		sin, _:=rd.ReadBytes('\n')
		lenv,_:=strconv.Atoi(strings.Trim(strings.Trim(string(sin), "\n"), "\r"))
		s,_ := rd.ReadBytes('\n')
		s = []byte(strings.Trim(strings.Trim(string(s), "\n"), "\r"))
		m:=make([]int, lenv)
		for j:=0; j<lenv; j++{
			m[j], _= strconv.Atoi(string(s[j]))
		}
		r:=fn(m)
		rss:=""
		for j:=0; j < lenv; j++{
			rss += strconv.Itoa(r[j])
		}
		fmt.Printf("%s\n", rss)
	}
}

func fn(b []int) (res []int){
	res = make([]int, len(b))
	res[0] = 1
	pr:=res[0] + b[0]
	for i:=1; i<len(b); i++{
		if pr != (1 + b[i]){
			res[i] = 1
			pr = 1 + b[i]
		}else{
			res[i] = 0
			pr = b[i]
		}
	}

	return res
}
