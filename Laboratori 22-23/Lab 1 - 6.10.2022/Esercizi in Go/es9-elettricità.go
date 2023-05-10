/*
9 Elettricità
Vogliamo leggere una sequenza di N interi (almeno 3), che descrivono il consumo di elettricità in N giorni, e stampare i giorni in cui il consumo è stato inferiore rispetto sia al giorno prima che al giorno dopo. I giorni sono numerati a partire da 1.
*/

package main 

import (
	"fmt";
	"os";
	"bufio";
	"strconv";
)

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}

	var giorni []int;
	giorni = make([]int, 0);
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);

	defer file.Close();
	
	for scanner.Scan(){
		giorno := scanner.Text();
		consumo,_:= strconv.Atoi(giorno);
		giorni = append(giorni, consumo);
	}

	for i:= 1; i < len(giorni)-1; i++{
		if giorni[i] < giorni[i-1] && giorni[i] < giorni[i+1]{
			fmt.Printf("Giorno %d consumo %d\nGiorno prima: %d\nGiorno dopo: %d\n\n", i+1, giorni[i], giorni[i-1], giorni[i+1]);
		}
	}

}