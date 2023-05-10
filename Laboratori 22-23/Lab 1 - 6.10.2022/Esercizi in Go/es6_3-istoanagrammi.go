/* 
3. Due parole costituiscono un anagramma se l’una si ottiene dall’altra permutando le lettere (es: attore, teatro). Scrivete un programma che legga due parole e verifichi se sono anagrammi.
*/


package main

import (
	"fmt";
	"os";
)

func isLetter(a byte) bool{
	if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z'{
		return true;
	}else{
		return false;
	}
}

func isUpper(a byte) bool{
	return a >= 'A' && a <= 'Z';
}

func main(){
	parola1:= os.Args[1];
	parola2:= os.Args[2];

	if len(parola1) != len(parola2){
		fmt.Println("Le due parole non sono anagrammi");
	}else{
		occorrenze1 := make(map[byte]int);
		occorrenze2 := make(map[byte]int);

		for i:=0; i < len(parola1); i++{
			if isLetter(parola1[i]){
				if isUpper(parola1[i]){
					occorrenze1[parola1[i]]++;
				}else{
					occorrenze1[parola1[i]-32]++;
				}
			}
		}

		for i:=0; i < len(parola2); i++{
			if isLetter(parola2[i]){
				if isUpper(parola2[i]){
					occorrenze2[parola2[i]]++;
				}else{
					occorrenze2[parola2[i]-32]++;
				}
			}
		}

		anag := true;
		for i,_ := range occorrenze1{
			if occorrenze1[i] != occorrenze2[i]{
				anag = false;
				break;
			}
		}
		if anag{
			fmt.Println("Le due parole sono anagrammi");
		}else{
			fmt.Println("Le due parole NON sono anagrammi");
		}
	}
}