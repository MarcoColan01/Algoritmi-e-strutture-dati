/*
7 Prime vocali
Si vuole scrivere un programma che legga una sequenza di stringhe e stampi, per ogni stringa, la sua prima vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole). Nel caso in cui una stringa non contenga vocali il programma stampa “la parola non contiene vocali”.
*/

package main 

import (
	"fmt";
	"bufio";
	"os";
)

func isVocale(c byte) bool {
	return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U';
}

func primaVocale(parola string){
	var voc byte = ' ';
	for i:=0; i < len(parola); i++ {
		if isVocale(parola[i]){
			voc = parola[i];
			break;
		}
	}
	if voc != ' '{
		fmt.Printf(" %c\n", voc);
	}else{
		fmt.Print(" La parola non contiene vocali\n");
	}
	return;
}

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);

	if err != nil{
		os.Exit(-1);
	}

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);
	defer file.Close();

	for scanner.Scan(){
		parola := scanner.Text();
		fmt.Printf("Prima vocale di %s:", parola);
		primaVocale(parola);
	}
}