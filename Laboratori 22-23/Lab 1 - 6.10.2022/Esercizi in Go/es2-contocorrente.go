/*
2 Conto corrente
Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un conto corrente. Il programma legge poi da standard input una serie di numeri interi che rappresentano spese da addebitare sul conto e stampa il saldo finale. La lettura si interrompe quando il saldo è <=0.
*/

package main

import(
	"fmt";
	"os";
	"strconv";
)

func main(){
	saldo,_:= strconv.Atoi(os.Args[1]);
	var addebito int;

	fmt.Println("Saldo iniziale:", saldo);
	for{
		fmt.Scanf("%d", &addebito);
		saldo += addebito;
		if saldo <= 0{
			break;
		}
	}

	fmt.Println("Saldo finale:", saldo);
}