/*
3 Andamento
Data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il nuovo valore è maggiore o uguale al precedente e ’-’ altrimenti.
*/

package main 

import (
	"fmt";
)

func main(){
	var sequenza []int;
	var n int;
	sequenza = make([]int, 0);

	for{
		fmt.Scanf("%d", &n);
		if n == 0{
			break;
		}
		sequenza = append(sequenza, n);
	}

	for i := 1; i < len(sequenza); i++{
		if sequenza[i] >= sequenza[i-1]{
			fmt.Print("+");
		}else{
			fmt.Print("-");
		}
	}
	fmt.Println();
}