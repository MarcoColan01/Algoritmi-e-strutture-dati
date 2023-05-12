package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Pila implementata come slice di interi dove pila[len(pila)-1] = cima della pila
func valuta(str string) int {
	pila := make([]int, 0)
	for i := 0; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			n, _ := strconv.Atoi(string(str[i]))
			pila = append(pila, n) //Push sulla pila
		} else if str[i] != ' ' {
			n1, n2 := pila[len(pila)-1], pila[len(pila)-2] //Pop sulla pila x2
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
