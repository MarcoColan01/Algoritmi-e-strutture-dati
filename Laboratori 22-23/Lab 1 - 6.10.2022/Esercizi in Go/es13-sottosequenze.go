/*
13 Somme delle sottosequenze crescenti
Scrivere un programma che legge una sequenza di interi e stampa la somma degli interi in ciascuna sottosequenza crescente. Il programma interrompe la lettura quando riceve un segnale di EOF. Ad esempio, su input 1 2 13 0 7 8 9 -1 0 2 il programma stampa le somme 16, 24 e 1.
*/

package main

import (
	"fmt";
	"os";
	"bufio";
	"strconv";
)

func stampaSottosomme(sequenza []int){
	sum := 0;
	for i:= 0; i < len(sequenza); i++{
		if i != len(sequenza)-1{
			if sequenza[i+1] <= sequenza[i]{
				sum += sequenza[i];
				fmt.Print(sum, ",");
				sum = 0;
			}else{
				sum += sequenza[i];
			}
		}else{
			sum += sequenza[i];
			fmt.Print(sum);
		}
	}
	fmt.Println();
	return;
}

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}

	scanner := bufio.NewScanner(file);
	var sequenza []int;
	sequenza = make([]int, 0);

	scanner.Split(bufio.ScanWords);
	defer file.Close();

	for scanner.Scan(){
		n,_:= strconv.Atoi(scanner.Text());
		sequenza = append(sequenza, n);
	}

	stampaSottosomme(sequenza);
}