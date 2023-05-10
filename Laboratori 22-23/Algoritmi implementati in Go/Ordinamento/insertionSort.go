package main

import (
	"fmt";
	"bufio";
	"os";
	"strconv";
)

func main(){
	filename := os.Args[1];
	file,err := os.Open(filename);
	if err != nil{
		os.Exit(-1);
	}

	scanner := bufio.NewScanner(file);
	var vet []int;
	vet = make([]int, 0);
	scanner.Split(bufio.ScanWords);
	defer file.Close();

	for scanner.Scan(){
		n,_:= strconv.Atoi(scanner.Text());
		vet = append(vet, n);
	}

	fmt.Print("Prima: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();

	//insertionSort
	//j:=0;
	for k:=1; k < len(vet); k++{
		x:= vet[k];
		j:= k-1;
		for{
			if !(j >= 0 && vet[j] > x){
				break
			}
			vet[j+1] = vet[j];
			j = j -1;
		}
		vet[j+1] = x;
	}
	fmt.Print("Dopo: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();
}