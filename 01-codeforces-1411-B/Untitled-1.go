package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ti:=scanner.Text()
	t , _:= strconv.Atoi(ti)
	for i:=0; i<t; i++{
		scanner.Scan()
		f:=scanner.Text()
		it,_:= strconv.ParseUint(f,10,64)
		fmt.Println(fn(it))
	}
}

func fn(n uint64) uint64   {
	if n < 10 {
		return n
	}

	i:=n
	for notgood(i){
		i++
	}

	return i

}

func notgood(nn uint64)  bool{
	n := nn
	var cache [10]bool
	for {
		m := n%10
		if !cache[m] && m!= 0 && m != 1 && nn % m != 0  {
			return true
		}
		cache[m] = true
		n = n / 10
		if n == 0{
			break
		}
	}

	return false
}
