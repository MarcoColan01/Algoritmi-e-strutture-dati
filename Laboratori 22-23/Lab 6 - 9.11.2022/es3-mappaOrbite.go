/*
La sfida di Advent of Code, 2019, day 6, intitolata Universal Orbit Map, può essere affrontata
modellando la situazione con un albero, e implementando l’albero con una tabella dei padri. La
prima parte della sfida è riassunta qui sotto.
Sei venuto in possesso della Mappa delle Orbite dell’Universo. La Mappa mostra che ogni
oggetto dello spazio (tranne il COM - Center of Mass) ruota attorno a esattamente un altro
oggetto. Questa relazione è indicata in questa forma: AAA)BBB significa che BBB orbita attorno
ad AAA.

*/

package main 
import(
	"fmt"
	"bufio"
	"strings"
	"os"
)

func passiPerSan(orbite map[string]string) int{
	cont := 0
}

func calcoloOrbite(orbite map[string]string) int{
	cont := 0 
	for _, padre := range orbite{
		app := padre 
		for app != ""{
			app = orbite[app]
			cont++
		}
	}
	return cont
}



func main(){
	filename := os.Args[1]
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	orbite := make(map[string]string)
	for scanner.Scan(){
		k := scanner.Text()
		s := strings.Split(k, ")")
		if s[1] == "COM"{
			orbite[s[1]] = ""
		}else{
			orbite[s[1]] = s[0]
		}
	}
	

	for figlio,padre := range orbite{
		fmt.Println(figlio,")",padre)
	}

	n := calcoloOrbite(orbite)
	fmt.Println("Numero totale orbite dirette e indirette:", n)

	n := passiPerSan(orbite)
	fmt.Println("Trasferimenti orbitali per la tratta YOU-SAN:", n)
}
