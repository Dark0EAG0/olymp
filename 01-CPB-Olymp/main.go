package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	n, _ := strconv.Atoi(scan.Text())
	scan.Scan()
	m, _ := strconv.Atoi(scan.Text())
	fmt.Println(fn(n, m))
}

func fn(n int, m int) (res int) {
	if (m+n)%2 != 0 {
		res++
		if n < m {
			m--
		} else {
			n--
		}
	}

	if n <= m {
		res += n * 2
	} else {
		res += m * 2
	}
	return res
}
