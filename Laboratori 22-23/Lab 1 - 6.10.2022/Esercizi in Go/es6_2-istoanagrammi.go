/*
2. Scrivete un programma che legga una riga di caratteri terminata da a-capo e che visualizzi un istogramma con una barra per ogni carattere dell’alfabeto, lunga quanto il numero delle sue occorrenze. Il programma non deve visualizzare le barre delle lettere che non compaiono e non deve fare distinzione fra maiuscole e minuscole (a tal fine potete usare le funzioni dichiarate nel file ctype.h).
*/

package main

import (
	"fmt";
	"bufio";
	"os";
	"strings"
)

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}
	occorrenze := make(map[string]int)

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanRunes);
	for scanner.Scan(){
		riga := scanner.Text();
		riga = strings.ToUpper(riga);
		if riga >= "A" && riga <= "Z"{
			occorrenze[riga]++;
		}
	}
	

	for i,j := range occorrenze{
		fmt.Print(i,": ");
		for k:=0; k < j; k++{
			fmt.Print("*");
		}
		fmt.Println();
	}

}