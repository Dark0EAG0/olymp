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
		n, _ := strconv.Atoi(las[0])
		k, _ := strconv.Atoi(las[1])
		str, _ := a.ReadString('\n')
		str = strings.Trim(strings.Trim(str, "\n"), "\r")
		mstr := make([]int64, n)
		mstra := strings.Split(str, " ")
		for j := 0; j < n; j++ {
			mstr[j], _ = strconv.ParseInt(mstra[j], 10, 64)
		}
		fmt.Printf("%d\n", fn(int64(k), mstr))
	}
}

func fn(k int64, a []int64) int64 {
	x := int64(0)
	sum := a[0]
	for i := 1; i < len(a); i++ {
		pr := (100*a[i] - k*sum + k - 1) / k
		x = max(x, pr)
		sum += a[i]
	}
	return x
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
