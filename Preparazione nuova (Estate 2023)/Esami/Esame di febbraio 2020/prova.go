package main

import "fmt"

func main() {
	a := []int{1, 4, 7, 8, 9, 10, 2, 5}
	a = a[0:0]
	fmt.Println(a)
}
