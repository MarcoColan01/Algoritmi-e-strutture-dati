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

	//integerSort
	var y []int;
	y = make([]int, 0);

	for i:=0; i < max+1; i++{
		y = append(y, 0);
	}

	for i:=0; i < len(vet); i++{
		x:=vet[i];
		y[x]++;
	}

	j:=0;
	for i:=0; i <= max; i++{
		for h:=0; h < y[i]; h++{
			vet[j]=i;
			j++;
		}
	}

	/*fmt.Print("Dopo: ");
	for i:=0; i < len(y); i++{
		fmt.Println(i, y[i]);
	}
	fmt.Println();*/
	fmt.Print("Prima: ");
	for i:=0; i < len(vet); i++{
		fmt.Print(vet[i], " ");
	}
	fmt.Println();

}