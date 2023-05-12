package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func valuta(str string) int {
	pila := make([]int, 0)
	for i := 0; i < len(str); i++ {
		if unicode.isDigit(str[i]) {
			n, _ := strconv.Atoi(str[i])
			pila = append(pila, n)
		} else {
			n1, n2 := pila[len(pila)-1], pila[len(pila)-2]
			pila = pila[:len(pila)-2]
			switch str[i] {
			case '+':
				pila = append(pila, n2+n1)
				break
			case '-':
				pila = append(pila, n2-n1)
				break
			case '/':
				pila = append(pila, n2/n1)
				break
			case '*':
				pila = append(pila, n2*n1)
				break
			default:
				break
			}
		}
	}
	return pila[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	n := valuta(str)
	fmt.Printf("%s -> %d\n", str, n)

}
