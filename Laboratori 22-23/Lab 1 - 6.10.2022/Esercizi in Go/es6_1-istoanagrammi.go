/*
6 Istogramma a anagrammi
1. Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga un’altra lettera e stampi quante volte la lettera compare nella prima riga.
*/

package main

import(
	"fmt";
	"os";
)

func main(){
	parola := os.Args[1];
	char := os.Args[2];

	cont := 0;

	for i:= 0; i < len(parola); i++{
		if parola[i] == char[0]{
			cont++;
		}
	}

	fmt.Println("il carattere", char, "compare", cont, "volte nella parola", parola);
}

