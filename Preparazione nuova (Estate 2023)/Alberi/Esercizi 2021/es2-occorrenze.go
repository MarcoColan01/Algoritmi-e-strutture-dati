package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type occorrenza struct {
	nome   string
	n      int
	dx, sx *occorrenza
}

type wordTree struct {
	root *occorrenza
}

func insertTree(r *occorrenza, nome string) *occorrenza {
	if r == nil {
		r = &occorrenza{nome, 1, nil, nil}
	} else if r.nome == nome {
		r.n++
	} else if nome < r.nome {
		r.sx = insertTree(r.sx, nome)
	} else {
		r.dx = insertTree(r.dx, nome)
	}
	return r
}

func visitInOrder(r *occorrenza) {
	if r == nil {
		return
	}
	visitInOrder(r.sx)
	fmt.Printf("%s %d\n", r.nome, r.n)
	visitInOrder(r.dx)
}

func searchWord(r *occorrenza, nome string) *occorrenza {
	for r != nil && r.nome != nome {
		if nome < r.nome {
			r = r.sx
		} else {
			r = r.dx
		}
	}
	return r
}

func convertiParola(nome string) string {
	if nome != "STOP" {
		nome = strings.ToLower(nome)
		if !unicode.IsLetter(rune(nome[len(nome)-1])) {
			nome = nome[:len(nome)-1]
		}
	}
	return nome
}

func main() {
	r := new(wordTree)
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	defer file.Close()

	cont := 0
	for scanner.Scan() {
		k := convertiParola(scanner.Text())
		if cont == 0 {
			if k != "STOP" {
				r.root = insertTree(r.root, k)
			} else {
				fmt.Println("OCCORRENZE DI PAROLE NEL TESTO IN ORDINE ALFABETICO:")
				visitInOrder(r.root)
				cont++
				fmt.Println("\nRICERCA DI PAROLE")
			}
		} else if cont == 2 {
			break
		} else if cont == 1 {
			app := searchWord(r.root, k)
			if app != nil && k != "STOP" {
				fmt.Printf("%s %d\n", k, 0)
			} else if k != "STOP" {
				fmt.Printf("%s %d", k, 0)
			}
		}
	}
}
