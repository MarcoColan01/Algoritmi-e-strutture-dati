/*
10 Coppie di numeri uguali
Scrivere un programma che legga da standard input una sequenza di numeri terminata da zero e conti quante sono le coppie di numeri naturali consecutivi uguali. 
*/

package main

import(
	"fmt";
)

func main(){
	var numeri []int;
	n := 0;
	cont := 0;
	numeri = make([]int,0);

	for{
		fmt.Scanf("%d", &n);
		if n == 0{
			break;
		}
		numeri = append(numeri, n);
	}

	for i:=0; i < len(numeri)-1; i++{
		if numeri[i] == numeri [i+1]{
			cont++;
		}
	}

	fmt.Println("Coppie di numeri consecutivi uguali:", cont);
}