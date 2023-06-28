package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vertex struct {
	c1, c2 int
}

type edge struct {
	e1, e2 vertex
}

type graph struct {
	v []vertex
	e map[edge]int
}

type app struct {
	e edge
	w int
}

func distanza(e1, e2 vertex) int {
	d := math.Abs(float64(e1.c1-e2.c1)) + math.Abs(float64(e1.c2-e2.c2))

	return int(d)
}

func costruisciPiano() graph {
	var g graph
	g.v = make([]vertex, 0)
	g.e = make(map[edge]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		if len(k) < 2 {
			break
		}
		n1, _ := strconv.Atoi(k[0])
		n2, _ := strconv.Atoi(k[1])
		var v vertex
		v.c1 = n1
		v.c2 = n2
		g.v = append(g.v, v)
	}

	for i := 0; i < len(g.v); i++ {
		for j := 0; j < len(g.v); j++ {
			if i != j {
				var e edge
				e.e1 = g.v[i]
				e.e2 = g.v[j]
				g.e[e] = distanza(g.v[i], g.v[j])
			}
		}
	}

	return g
}

func find(v1 vertex, parent map[vertex]vertex) vertex {
	for parent[v1] != v1 {
		v1 = parent[v1]
	}
	return v1
}

func union(parent map[vertex]vertex, i, j vertex) {
	a := find(i, parent)
	b := find(j, parent)
	parent[a] = b
}

func calcolaLink(g graph) int {
	mincost := 0
	parent := make(map[vertex]vertex)

	for i := 0; i < len(g.v); i++ {
		parent[g.v[i]] = g.v[i]
	}

	pesi := make([]app, 0)

	for i, k := range g.e {
		var j app
		j.e = i
		j.w = k
		pesi = append(pesi, j)
	}

	sort.SliceStable(pesi, func(i, j int) bool { return pesi[i].w < pesi[j].w })

	edgeCount := 0

	for edgeCount < len(g.e) {
		a := find(pesi[edgeCount].e.e1, parent)
		b := find(pesi[edgeCount].e.e2, parent)

		if a != b {
			union(parent, a, b)
			mincost += pesi[edgeCount].w
			fmt.Println(a, b)
		}
		edgeCount++
	}

	return mincost
}

func main() {
	g := costruisciPiano()
	n := calcolaLink(g)
	fmt.Println(n)
}
