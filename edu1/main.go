package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rd:=bufio.NewScanner(os.Stdin)
	rd.Scan()
	l,_:=strconv.Atoi(rd.Text())
	for i:=0; i< l ; i++{
		rd.Scan()
		fmt.Println(fn(rd.Text()))
	}
}

func fn(s string) (res int){
	i, j :=0,0
	for ; i<len(s)-1; i++{
		j = i+1
		bt:=1
		str:=s[i:j]
		for ; j<(-(i-(len(s)-1))); j++{
			if comp(str){
				bt= j-i+1
			}
		}
		if res < bt{
			res = bt
		}
	}
	return res
}

func comp(s1 string) bool {
	for i:=0;i<(len(s1)/2); i++{
		if s1[i] != s1[(len(s1) - 1)  - i ]{
			return false
		}
	}
	return true
}