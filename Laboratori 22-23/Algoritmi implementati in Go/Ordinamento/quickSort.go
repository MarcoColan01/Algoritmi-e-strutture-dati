package main

import (
	"fmt";
	"bufio";
	"os";
	"strconv";
)

func partiziona(vet []int, i int, f int) int{
	perno:=vet[i];
	sx := i; 
	dx := f;
	for {
		if !(sx < dx){
			break;
		}
		for{
			dx--;
			if !(vet[dx] > perno){
				break;
			}
		}
		for{
			sx++;
			if !(vet[sx] <= perno && sx < dx){
				break;
			}
		}
		if sx < dx{
			app := vet[sx];
			vet[sx]=vet[dx];
			vet[dx]=app;
		}
	}
	app := vet[i];
	vet[i]=vet[dx];
	vet[dx]=app;
	return dx;
}

func quickSort(vet []int, i int, f int) []int{
	for{
		if !(f-i > 1){
			break;
		}
		m:=partiziona(vet,i,f);
		if m-i < f-m{
			quickSort(vet,i,m);
			i=m+1;
		}else{
			quickSort(vet, m+1, f);
			f=m;
		}
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

//quickSort
vet=quickSort(vet, 0, len(vet));

fmt.Print("Dopo: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();
}