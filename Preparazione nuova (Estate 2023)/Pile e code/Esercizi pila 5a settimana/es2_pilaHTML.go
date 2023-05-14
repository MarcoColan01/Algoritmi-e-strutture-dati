package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkHTML(str string) {
	pila := make([]string, 0)
	var i int
	k := strings.Split(str, " ")
	for i = 0; i < len(k); i++ {
		if k[i][1] != '/' {
			pila = append(pila, k[i])
		} else {
			t2 := pila[len(pila)-1]
			pila = pila[:len(pila)-1]
			if t2[1] != k[i][2] {
				break
			}
		}
	}
	if i < len(k) {
		fmt.Printf("errore in pos %d\n", i+1)
	} else if len(pila) > 0 {
		fmt.Printf("sono rimasti aperti i seguenti tag: %s\n", pila)
	} else {
		fmt.Printf("documento HTML ben formattato\n")
	}
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan() {
		k := scanner.Text()
		checkHTML(k)
	}
}
