package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
	Modellazione
	Nodi: stati
	Archi: transizioni da uno stato all'altro
	Delta transizione: se x è la posizione, !x, !x-1, !x+1

	Strategia: ricerca in ampiezza
*/

type state struct {
	config []bool
}

type trackingState struct {
	state    state
	distance int
}

func main() {
	startingState := parseInput(os.Args[1])
	finalState := parseInput(os.Args[2])
	path, length := breadthFirstSearch(startingState, finalState)
	printPath(path)
	fmt.Println("Path length:", length)
}

func parseInput(input string) (state state) {
	for _, value := range input {
		temp, _ := strconv.Atoi(string(value))
		if temp == 1 {
			state.config = append(state.config, true)
		} else {
			state.config = append(state.config, false)
		}
	}
	return
}

func transition(center int, from state) (next state) {
	if center == 0 {
		for index, value := range from.config {
			if index == 0 || index == 1 || index == len(from.config)-1 {
				next.config = append(next.config, !value)
			} else {
				next.config = append(next.config, value)
			}
		}
	} else if center == len(from.config)-1 {
		for index, value := range from.config {
			if index == 0 || index == center || index == center-1 {
				next.config = append(next.config, !value)
			} else {
				next.config = append(next.config, value)
			}
		}
	} else {
		for index, value := range from.config {
			if index == center || index == center-1 || index == center+1 {
				next.config = append(next.config, !value)
			} else {
				next.config = append(next.config, value)
			}
		}
	}
	return
}

func breadthFirstSearch(from, to state) ([]string, int) {
	var seen map[*state]bool = map[*state]bool{}
	queue := []trackingState{}
	queue = append(queue, trackingState{from, 0})
	seen[&from] = true
	var parent map[string]string = make(map[string]string)
	parent[from.toString()] = ""

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		for index := range s.state.config {
			next := transition(index, s.state)
			parent[next.toString()] = s.state.toString()

			if next.equals(to) {

				/*
					L'idea sarebbe tenere una mappa dei padri da puntatore a stato a puntatore a stato.
					Non mi piace tanto fare una map[string]string dopo che ho fatto di tutto per non
					tenere il grafo in memoria, ma fare una mappa da *state a *state è un incubo.
					Sembra ragionevolmente piccola comunque (Se il grafo fosse completamente connesso dovremmo avere
					una mappa con V^2 chiavi al più?).
				*/

				// Ricostruzione sequenza
				path := []string{}
				current := next.toString()

				for i := 0; i < s.distance+1; i++ {
					path = append(path, current)
					current = parent[current]
				}
				// Fine ricostruzione sequenza
				fmt.Println(parent)
				return path, s.distance + 1
			} else if !seen[&next] {
				queue = append(queue, trackingState{next, s.distance + 1})
				seen[&next] = true
			}
		}
	}
	return []string{}, math.MaxInt
}

// Si potrebbero confrontare anche le stringhe, ma anche equals() è venuto prima.
func (s state) equals(t state) bool {
	for index, value := range s.config {
		if t.config[index] != value {
			return false
		}
	}
	return true
}

func (s state) toString() (result string) {
	for _, value := range s.config {
		if value {
			result += "1"
		} else {
			result += "0"
		}
	}
	return
}

// Abbastanza inutile, ma è venuta prima di toString(), quindi l'ho lasciata per non dover riscrivere il codice.
func (s state) print() {
	fmt.Println(s.toString())
}

func printPath(path []string) {
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}
	fmt.Println()
}
