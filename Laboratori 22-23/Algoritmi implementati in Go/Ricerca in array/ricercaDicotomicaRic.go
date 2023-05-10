//ricerca dicotomica ricorsiva in array ordinato. Tempo proporzionale a log(n)

package main 
import(
	"fmt"
)

func ricercaDicRic(a []int, sx int, dx int, x int) int{
	if dx <= sx {
		return -1
	}else{
		m := (dx+sx)/2
		if x == a[m]{
			return m 
		}else if x < a[m]{
			return ricercaDicRic(a, sx, m, x)
		}else{
			return ricercaDicRic(a, m+1, dx, x)
		}
	}

}

func main(){
	a := []int{1,2,5,9,12,34,56,78,80,83,87,89,95,100,112,114,120,125,145,200,207,309,345,367,400}

	//k := ricercaDicRic(a, 0, len(a), 367)
	k := ricercaDicRic(a, 0, len(a), 50)

	if k < 0{
		fmt.Println("Elemento non trovato");
	}else{
		fmt.Println("Elemento trovato in posizione", k);
	}
}