package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func converti(str string) string {
	var conv string
	pila := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			if unicode.IsDigit(rune(str[i])) {
				conv = conv + string(str[i]) + " "
			} else if str[i] == ')' {
				conv = conv + string(pila[len(pila)-1]) + " "
				pila = pila[:len(pila)-1]
			} else if isOperator(rune(str[i])) {
				pila = append(pila, str[i])
			}
		}
	}
	return conv
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	str := scanner.Text()
	conv := converti(str)
	fmt.Printf("%s -> %s\n", str, conv)
}
