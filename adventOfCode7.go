package main

import (
	"bufio"
	"os"
	"strings"
)

type elems struct {
	str   string
	size  int
	padre *elems
	figli []*elems
}

type filesystem struct {
	system map[string]elems
}

func main() {
	var fs filesystem
	var app elems
	fs.system = make(map[string]elems)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan() {
		f := ""
		k := strings.Split(scanner.Text(), " ")
		if k[0] == "$" {
			if k[2] != ".." {
				f = k[2]
				app.str = f
				if f == "/" {
					app.padre = nil
				} else {

				}
			}
		}
	}
}
