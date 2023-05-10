/*
Dei sassi sferici sono ammucchiati a formare una piramide, con un sasso in cima, posto al centro di un quadrato formato da 4 sassi (2 per lato), posti a loro volta sopra un quadrato formato da 9 sassi (3 per lato), e così via.
Scrivete una funzione ricorsiva sassi(height int) int che, data l’altezza height della piramide, restituisca il numero di sassi che la compongono.
Ad esempio: f(1) deve restituire 1 e f(2) deve restituire 5.  */


package main 

import(
	"os"
	"strconv"
	"fmt"
)

func sassi(h int) int{
	if h == 1{
		return 1
	}

	return h*h + sassi(h-1);
}

func main(){
	h,_ := strconv.Atoi(os.Args[1]);
	fmt.Println("Il numero di sassi della piramide di altezza",h,"è:", sassi(h));
}

