/*
5 Elaborazione di stringhe
Scrivete le funzioni specificate sotto.

1. Scrivere una funzione quanteConA(parole []string) int che, data una slice di stringhe, restituisca quante stringhe contengono il carattere ‘a’.

2. Scrivere una funzione primaConA(parole []string) string che, data una slice di stringhe, restituisca la prima parola che contentiene il carattere ‘a?, o la stringa vuota.

3. Scrivere una funzione piuCorta(parole []string) int che, data una slice di stringhe, restituisca la lunghezza della stringa più corta in termini di byte.

*/

package main

import (
	"fmt";
	"bufio";
	"os";
)

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}
	var parole []string;
	parole = make([]string, 0);

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);
	for scanner.Scan(){
		riga := scanner.Text();
		parole = append(parole, riga);
	}

	fmt.Println("Parole che contengono la a: ", quanteConA(parole));

	parola := primaConA(parole);
	if parola != ""{
		fmt.Println("La prima parola che contiene a? è: ", parola);
	}else{
		fmt.Println("Nessuna parola contiene a?");
	}
}

func piuCorta(string []parole) int{
	
}

func contieneA(parola string) bool{
	trovata := false;
	for i := 0; i < len(parola); i++{
		if parola[i] == 'a'{
			trovata = true;
			break;
		}
	}
	return trovata;
}

func quanteConA(parole []string) int{
	cont := 0;
	for i:=0; i < len(parole); i++{
		if contieneA(parole[i]){
			cont++;
		}
	}
	return cont;
}

func contieneaqst(parola string) bool{
	trovata := false;

	for i:= 0; i < len(parola)-1; i++{
		if parola[i] == 'a' && parola[i+1] == '?'{
			trovata = true;
			break;
		}
	}

	return trovata;
}

func primaConA(parole []string) string{
	parola := "";
	for i:=0; i < len(parole); i++{
		if contieneaqst(parole[i]){
			parola = parole[i];
			break;
		} 
	}
	return parola;
}