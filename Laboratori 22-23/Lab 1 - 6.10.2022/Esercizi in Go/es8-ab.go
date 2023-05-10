/*
8 Inizia con ’a’, inizia con ’b’
Si vuole scrivere un programma che legga una sequenza di dieci stringhe e stampi il numero di stringhe che cominciano con la lettera a e il numero di stringhe che cominciano con b. Per calcolare il numero di stringhe che cominciano per a, usiamo il piano del conteggio; lo stesso piano serve anche a calcolare il numero di stringhe che cominciano per b.
*/

package main 

import(
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

	i := 0;
	numA := 0;
	numB := 0;
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);

	for scanner.Scan(){
		parola := scanner.Text();
		if parola[0] == 'a' || parola[0] == 'A'{
			numA++;
		} else if parola[0] == 'b' || parola[0] == 'B'{
			numB++;
		}
		i++;
		if i == 10{
			break;
		}
	}
	fmt.Printf("Numero di parole che iniziano per a: %d\nNumero di parola che iniziano per b: %d\n", numA, numB);
}