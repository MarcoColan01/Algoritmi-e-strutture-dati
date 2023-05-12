package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var str string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	str := strings.Split(scanner.Text(), " ")
	//k := strings.Split(str, " ")
	fmt.Println(str)

}
