/*
12 Massimo numero di zeri
Scrivere un programma massimo_n_zeri.go che legge da standard input una sequenza di interi terminata da -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi siano più numeri che contengono il massimo numero di 0, il programma stampa l?ultimo incontrato. Ad esempio, se la sequenza letta è 3040145 80 1707 105002 78 1970 6005 -1 il programma stampa 105002.
*/

package main 

import(
	"fmt";
	"bufio";
	"os";
)

func numZeri(num string) int{
	cont := 0;
	for i:=0; i < len(num); i++{
		if num[i] == '0'{
			cont++;
		}
	}
	return cont;
}

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}
	max := "0";
	maxZeri := 0;

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);
	defer file.Close();

	for scanner.Scan(){
		num := scanner.Text();
		if num == "-1"{
			break;
		}
		k := numZeri(num);

		if k >= maxZeri{
			maxZeri = k;
			max = num;
		}
	}

	fmt.Println("Il numero della sequenza con più 0 è", max, "con", maxZeri, "zeri.");
}
