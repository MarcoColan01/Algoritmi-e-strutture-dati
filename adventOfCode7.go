package main

import (
	"bufio"
	"os"
)

type elems struct {
	str   string
	size  int
	padre *elems
	figli []*elems
}

type filesystem struct {
	vet []elems
}

func main() {
	var fs filesystem
	fs.vet = make([]elems, 0)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan() {
		k := scanner.Text()
	}
}
