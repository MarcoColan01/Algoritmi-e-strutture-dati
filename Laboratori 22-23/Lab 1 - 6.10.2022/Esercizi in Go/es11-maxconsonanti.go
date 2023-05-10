/*
11 Massimo numero vocali
Si vuole scrivere una funzione che, presa come argomento una slice di stringhe, restituisca il numero massimo di consonanti contenute in una stringa (assumiamo per semplicità che le stringhe contengano solo caratteri minuscoli)
*/

package main 

import(
	"fmt";
	"bufio";
	"os";
	"strings"
)

func isLetter(c byte) bool{
	return c >= 'a' && c <= 'z';
}

func isVocal(c byte) bool{
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
}

func nConsonanti(parola string) int{
	cont := 0;
	parola = strings.ToLower(parola);
	for i:=0; i < len(parola); i++{
		if isLetter(parola[i]) && !isVocal(parola[i]){
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

	scanner := bufio.NewScanner(file);
	max := 0;
	scanner.Split(bufio.ScanWords);
	defer file.Close();

	for scanner.Scan(){
		parola := scanner.Text();
		k := nConsonanti(parola);
		if k > max{
			max = k;
		}
	}

	fmt.Println("La parola con più consonanti ne ha", max);
}