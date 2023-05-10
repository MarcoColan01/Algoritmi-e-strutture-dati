//ricerca sequenziale in un array non ordinato. Tempo proporzionale ad n

package main
import(
	"fmt"
)

func ricerca(a []int, k int) int{
	i := len(a)-1
	for{
		if !(i >= 0 && a[i] != k){
			break
		}
		i--
	}
	return i
}

func main(){
	a := []int{10, -5, 11, 23, 45, 9, 21, 2, 7, 18, 29, 1, 33, 100, 27, 31, 9, 87, 75, 65, 12, 99, 34, 67, 8, 4, 37, 60}

	//k := ricerca(a, 31)
	k := ricerca(a,6)

	if k < 0{
		fmt.Println("Elemento non trovato");
	}else{
		fmt.Println("Elemento trovato in posizione", k);
	}
}