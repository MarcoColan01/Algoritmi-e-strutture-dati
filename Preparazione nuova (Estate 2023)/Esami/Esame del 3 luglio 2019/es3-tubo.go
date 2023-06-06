package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func segmentiMax(a []int, l int, n int) int {
	cont, somma := 0, 0
	sort.Ints(a)
	for i := 0; i < n; i++ {
		if somma+a[i] <= l {
			somma += a[i]
			cont++
		}
	}
	return cont
}

func main() {
	var l, n int
	a := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		if len(k) == 2 {
			l, _ = strconv.Atoi(k[0])
			n, _ = strconv.Atoi(k[1])
		} else if len(k) > 2 {
			for i := 0; i < len(k); i++ {
				n, _ := strconv.Atoi(k[i])
				a = append(a, n)
			}
		} else {
			break
		}
	}
	max := segmentiMax(a, l, n)
	fmt.Println(max)
}
