package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcoloOrbite(orbite map[string]string) int {
	cont := 0
	var app string
	for figlio, _ := range orbite {
		app = figlio
		for {
			if orbite[app] == "NIL" {
				break
			}
			app = orbite[app]
			cont++
		}
	}
	return cont
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	orbite := make(map[string]string)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), ")")
		orbite[k[1]] = k[0]
	}

	n := calcoloOrbite(orbite)
	fmt.Printf("Numero orbite totali: %d\n", n)
}
