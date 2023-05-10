package main

import (
	"fmt";
	"bufio";
	"os";
	"strconv";
)

func merge(vet []int, i int, m int, f int, x []int) []int{
	i1 := i;
	i2 := m;
	k := 0;

	for{
		if !(i1 < m && i2 < f){
			break;
		}

		if vet[i1] <= vet[i2]{
			x[k] = vet[i1];
			i1++;
		}else{
			x[k] = vet[i2];
			i2++;
		}
		k++;
	}

	if i1 < m{
		for j:=i1; j < m; j++{
			x[k] = vet[j];
			k++;
		}
	}else{
		for j:=i2; j < f; j++{
			x[k] = vet[j];
			k++;
		}
	}

	for k:=0; k < f-i; k++{
		vet[i+k] = x[k];
	}

	return vet;
}

func mergeSort(vet []int, i int, f int, x []int) []int{
	if f-i > 1{
		m:= (i+f)/2;
		mergeSort(vet, i, m, x);
		mergeSort(vet, m, f, x);
		vet = merge(vet, i, m, f, x);
	}

	return vet;
}

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

	max := -1;
	for scanner.Scan(){
		n,_:= strconv.Atoi(scanner.Text());
		vet = append(vet, n);
		if n > max{
			max = n;
		}
	}
	
	fmt.Print("Prima: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();

	//mergeSort

	var x []int;
	x = make([]int, len(vet));
	vet = mergeSort(vet, 0, len(vet), x);

	fmt.Print("Dopo: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();

}