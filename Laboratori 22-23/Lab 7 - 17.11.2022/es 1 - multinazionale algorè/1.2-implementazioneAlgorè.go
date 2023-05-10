/*
Definite, in Go, uno o più tipi di dati utili a rappresentare i dipendenti e le loro relazioni, in base
alle scelte fatte ai punti precedenti.
Scrivete quindi una funzione Go per ciascuno dei compiti enunciati sopra (potete scrivere funzioni ausiliarie se lo ritenete opportuno). Le funzioni devono avere questi nomi:
	(a) stampaSubordinati
	(b) quantiSenzaSubordinati
	(c) supervisore
	(d) stampaImpiegatiSopra
	(e) stampaImpiegatiConSupervisore
Scrivete anche un programma per testare le funzioni.
*/

package main 
import(
	"fmt"
	"bufio"
	"strings"
	"os"
)

func stampaSubordinati(gerarchie map[string]string, dip string){
	fmt.Print("Subordinati per ",dip,": ")
	for i,k := range gerarchie{
		if k == dip{
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	return
}

func stampaImpiegatiConSupervisore(gerarchie map[string]string){
	for i,k := range gerarchie{
		if k != ""{
			fmt.Println(i,":",k)
		}
	}
	return
}

func quantiSenzaSubordinati(gerarchie map[string]string) int{
	sup := make(map[string]bool)
	for _,k := range gerarchie{
		if k != ""{
			if sup[k] == false{
				sup[k] = true
			}
		}
	}
	return len(gerarchie)-len(sup)
}

func supervisore(gerarchie map[string]string, dip string){
	fmt.Println("Supervisore di ",dip,": ", gerarchie[dip])
	return
}

func stampaImpiegatiSopra(gerarchie map[string]string, dip string){
	app := gerarchie[dip]
	fmt.Print("Impiegati sopra ", dip, ": ")
	for {
		fmt.Print(app, " ")
		app = gerarchie[app]
		if gerarchie[app] == ""{
			fmt.Print(app)
			break
		}
	}
	fmt.Println()
	return
}

func main(){
	filename := os.Args[1]
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}
	gerarchie := make(map[string]string)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan(){
		k := strings.Split(scanner.Text(), ":")
		gerarchie[k[0]] = k[1]
	}
	stampaImpiegatiConSupervisore(gerarchie)
	stampaSubordinati(gerarchie, "Anna")
	n := quantiSenzaSubordinati(gerarchie)
	fmt.Println(n)
	stampaImpiegatiSopra(gerarchie, "Daniela")
	supervisore(gerarchie, "Irene")
}
