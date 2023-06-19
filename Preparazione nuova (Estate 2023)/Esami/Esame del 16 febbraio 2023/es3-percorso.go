package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type lines struct {
	nLinea   int
	stazioni []string
}

type rete struct {
	corrispondenze map[string][]int
	linee          []lines
	adiacenti      map[string][]string
}

func leggiDati(nomeFile string) rete {

	var r rete
	r.linee = make([]lines, 0)
	r.adiacenti = make(map[string][]string)
	r.corrispondenze = make(map[string][]int)
	file, err := os.Open(nomeFile)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := scanner.Text()
		n, _ := strconv.Atoi(strings.Split(k, " ")[1][0:1])
		var l lines
		l.nLinea = n
		l.stazioni = make([]string, 0)
		app := strings.Split(k, ":")
		st := strings.Split(app[1], ",")
		st[len(st)-1] = strings.Trim(st[len(st)-1], ".")
		for i := 0; i < len(st); i++ {
			st[i] = strings.TrimSpace(st[i])
			l.stazioni = append(l.stazioni, st[i])
			r.corrispondenze[st[i]] = append(r.corrispondenze[st[i]], l.nLinea)
			if i == 0 {
				st[i+1] = strings.TrimSpace(st[i+1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			} else if i == len(st)-1 {
				st[i-1] = strings.TrimSpace(st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			} else {
				st[i+1] = strings.TrimSpace(st[i+1])
				st[i-1] = strings.TrimSpace(st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			}
		}
		r.linee = append(r.linee, l)
	}
	return r
}

func stessaLinea(mm rete, s1 string, s2 string) bool {
	stessa := false
	app1 := mm.corrispondenze[s1]
	app2 := mm.corrispondenze[s2]
	if len(app1) == 1 && len(app2) == 1 {
		if app1[0] == app2[0] {
			stessa = true
		}
	} else {
		for i := 0; i < len(app1); i++ {
			if stessa {
				break
			}
			for j := 0; j < len(app2); j++ {
				if app1[i] == app2[j] {
					stessa = true
					break
				}
			}
		}
	}
	return stessa
}

func percorsoMinimo(mm rete, partenza, arrivo string) []string {
	distanze := make(map[string]int)
	prev := make(map[string]string)
	coda := []string{partenza}
	for i, _ := range mm.corrispondenze {
		distanze[i] = math.MaxInt
	}
	distanze[partenza] = 0

	for len(coda) > 0 {
		app := coda[0]
		coda = coda[1:]
		if app == arrivo {
			break
		}
		for i := 0; i < len(mm.adiacenti[app]); i++ {
			if distanze[mm.adiacenti[app][i]] == math.MaxInt {
				if len(mm.corrispondenze[app]) == 1 {
					distanze[mm.adiacenti[app][i]] = distanze[app] + 1
				} else {
					if !stessaLinea(mm, mm.adiacenti[app][i], app) {
						distanze[mm.adiacenti[app][i]] = distanze[app] + 2
					} else {
						distanze[mm.adiacenti[app][i]] = distanze[app] + 1
					}
				}

				coda = append(coda, mm.adiacenti[app][i])
				prev[mm.adiacenti[app][i]] = app
			}
		}
		//fmt.Println(coda)

	}
	path := make([]string, 0)     
	for i, k := range prev {
		app := k
		for app != partenza {
			path = append(path, app)
			app = 
		}
	}
	return path
}

func tempo(mm rete, partenza, arrivo string) int {
	distanze := make(map[string]int)
	coda := []string{partenza}
	for i, _ := range mm.corrispondenze {
		distanze[i] = math.MaxInt
	}
	distanze[partenza] = 0

	for len(coda) > 0 {
		app := coda[0]
		coda = coda[1:]
		if app == arrivo {
			break
		}
		for i := 0; i < len(mm.adiacenti[app]); i++ {
			if distanze[mm.adiacenti[app][i]] == math.MaxInt {
				if len(mm.corrispondenze[app]) == 1 {
					distanze[mm.adiacenti[app][i]] = distanze[app] + 1
				} else {
					if !stessaLinea(mm, mm.adiacenti[app][i], app) {
						distanze[mm.adiacenti[app][i]] = distanze[app] + 2
					} else {
						distanze[mm.adiacenti[app][i]] = distanze[app] + 1
					}
				}

				coda = append(coda, mm.adiacenti[app][i])
			}
		}
		//fmt.Println(coda)

	}

	return distanze[arrivo]
}

func main() {
	r := leggiDati("input-metro.txt")

	n := tempo(r, "Moscova", "Marche")
	fmt.Println(n)

	k := percorsoMinimo(r, "Moscova", "Marche")
	fmt.Println(k)
	//stampaRete(r)
}
