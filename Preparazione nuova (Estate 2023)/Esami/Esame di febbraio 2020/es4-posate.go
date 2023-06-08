package main

import "fmt"

func robot() bool {
	var sequenza string
	fmt.Scan(&sequenza)
	var pila []rune

	sx := ' '
	dx := ' '

	for i := 0; i < len(sequenza); i++ {
		fmt.Printf("%c dx %c sx%c\n", pila, dx, sx)
		if dx == ' ' && sx == ' ' { //bracci liberi
			sx = rune(sequenza[i])
		} else if sx != ' ' && dx == ' ' { //braccio sx occupato ma dx libero
			if len(pila) > 0 {
				if pila[len(pila)-1] != sx {
					dx = pila[len(pila)-1]
					pila = pila[:len(pila)-1]
				}
			} else {
				dx = rune(sequenza[i])
			}
		}
		if sx != ' ' && dx != ' ' { //bracci occupati
			if sx == dx {
				pila = append(pila, dx)
				dx = ' '
			} else {
				dx = ' '
				sx = ' '
			}
		}
	}

	fmt.Printf("%c dx %c sx%c\n", pila, dx, sx)
	return len(pila) == 0 && dx == ' ' && sx == ' '
}

func main() {
	if robot() {
		fmt.Println("Successo")
	} else {
		fmt.Println("Fallimento")
	}
}
