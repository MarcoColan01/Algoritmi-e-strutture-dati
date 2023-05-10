/*
Ambientazione
Un pesce lanterna crea un nuovo pesce lanterna ogni 7 giorni. Ma il processo non è necessariamente sincronizzato, dipende dall’età dei pesci (un pesce lanterna di 5 giorni farà nascere un nuovo pesce fra 2 giorni; un
pesce che ha 3 giorni lo farà nascere fra 4 giorni).
Esempio
Se un pesce lanterna ha un timer con valore 3, allora:
    • Dopo un giorno, il timer sarà di 2
    • Dopo un altro giorno, il timer sarà di 1
    • Dopo un altro giorno, il timer sarà di 0
    • Dopo un altro giorno, il timer ricomincerà da 6, e sarà nato un nuovo pesce lanterna con timer con valore 8 (cioè 6+2)
    • Dopo un altro giorno, il primo pesce avrà timer con valore 5, e l’altro 7

NB: Notate che il pesce lantera riparte da 6, non da 7 (poiché 0 è incluso come valore valido del timer). Il pesce lanterna nuovo inizierà con timer a 8, che comincerà a decrescere dal giorno successivo.  */

package main
 
import(
	"fmt";
	"os";
	"bufio";
	"strconv";
)

func evoluzionePesci(pesci []int) []int{
	new := 8;
	len := len(pesci);
	for i:=0; i < len; i++{
		if pesci[i] == 0{
			pesci[i] = 6;
			pesci = append(pesci, new);
		}else{
			pesci[i]--;
		}
	}
	return pesci;
}

func main(){
	filename := os.Args[1];
	giorni,_ := strconv.Atoi(os.Args[2]);
	file,err := os.Open(filename);

	if err != nil{
		os.Exit(-1);
	}

	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanWords);

	defer file.Close();

	var pesci []int;
	pesci = make([]int, 0);

	for scanner.Scan(){
		n,_ := strconv.Atoi(scanner.Text());
		if n >= 0 && n <= 6{
			pesci = append(pesci, n);
		}
	}


	for i:=0; i < giorni; i++{
		pesci = evoluzionePesci(pesci);
		fmt.Print("Dopo il giorno ", i+1, " ");
		for j:=0; j < len(pesci); j++{
			fmt.Print(pesci[j], " ");
		}
		fmt.Println();
	}

	fmt.Println("Dopo", giorni, "giorni ci sono", len(pesci), "pesci")
}