// ricerca dicotomica iterativa in array ordinato. Tempo proporzionale a log(n), migliora la versione ricorsiva perchè non usa spazio aggiuntiuvo per lo stack della ricorsione. 

package main
import(
	"fmt"
)

func ricercaDicIter(a []int, x int) int{
	sx := 0
	dx := len(a)
	pos := -1

	for{
		if !(sx < dx && pos == -1){
			break
		}
		m := (sx+dx)/2
		if x == a[m]{
			pos = m
		}else if x < a[m]{
			dx = m
		}else{
			sx = m+1
		}
	}

	return pos
}

func main(){
	a := []int {1,2,5,9,12,34,56,78,80,83,87,89,95,100,112,114,120,125,145,200,207,309,345,367,400}

	//k := ricercaDicIter(a, 367)
	k := ricercaDicIter(a, 31)
	if k < 0{
		fmt.Println("Elemento non trovato");
	}else{
		fmt.Println("Elemento trovato in posizione", k);
	}
}