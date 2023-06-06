package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type reparto struct {
	prog, nLav    int
	dx, sx, padre int
}

type produzione struct {
	reparti map[int]*reparto
}

func costruisciReparti() *produzione {
	p := new(produzione)
	p.reparti = make(map[int]*reparto)
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), ":")
		r := strings.Split(k[1][1:], " ")

		n, _ := strconv.Atoi(k[0])
		app := &reparto{n, 0, 0, -1, -1}
		for i := 0; i < len(r); i++ {
			n, _ := strconv.Atoi(r[i])
			if i == 0 {
				app.nLav = n
			} else if i == 1 {
				app.sx = n
			} else {
				app.dx = n
			}
		}

	}
	return p

}

func main() {
	p := costruisciReparti()
	for i, k := range p.reparti {
		fmt.Printf("Reparto %d: lavoratori: %d reparti succesivi %d %d\n", i, k.nLav, k.sx, k.dx)
	}
}
