//GIUSTO

package main 
import(
	"fmt"
)

func stampaSlice(a []int){
	for i:=0; i < len(a); i++{
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	return
}

func risistema(a []int, r int, l int){
	v := r 
	x := a[v]
	daCollocare := true

	for{
		if (2*v +1) >= l {
			daCollocare = false
		}else{
			u := 2*v +1
			if u+1 < l && a[u+1] > a[u]{
				u++
			}
			if a[u] > x{
				a[v] = a[u]
				v = u
			}else{
				daCollocare = false
			}
		}
		if !daCollocare{
			break
		}
	}
	a[v] = x
}

func creaHeap(a []int){
	for i:=(len(a)/2); i >= 0; i--{
		risistema(a, i, len(a))
	}
}

func HeapSort(a []int){
	creaHeap(a)
	for l:= len(a)-1; l >= 1; l--{
		a[0], a[l] = a[l], a[0]
		risistema(a, 0, l)
	}
}

func main(){
	a := []int {10, -5, 11, 23, 45, 9, 21, 2, 7, 18, 29, 1, 33, 100, 27, 31, 9, 87, 75, 65, 12, 99, 34, 67, 8, 4, 37, 60}

	fmt.Print("Prima: ")
	stampaSlice(a)

	HeapSort(a)

	fmt.Print("Dopo: ")
	stampaSlice(a)
}