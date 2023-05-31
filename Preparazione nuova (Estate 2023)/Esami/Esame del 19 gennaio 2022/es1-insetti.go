package main

import "fmt"

//Complessità: O(n)
func contaAumenti(vet []int) int {
	cont := 0
	for i := 0; i < len(vet)-2; i++ {
		if vet[i] < vet[i+1] {
			cont++
		}
	}
	return cont
}

//Complessità: O(n)
func contaAumentiFin(vet []int, m int) int {
	cont := 0
	dyn := make([]int, len(vet))
	dyn[0] = vet[0]
	dyn[1] = vet[1] + vet[0]
	dyn[2] = dyn[1] + vet[2]
	for i := 3; i < len(vet); i++ {
		dyn[i] = vet[i] + vet[i-1] + vet[i-2]
		if dyn[i] > dyn[i-1] {
			cont++
		}
	}
	return cont
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	rilevazioni := []int{5689, 6258, 4923, 3926, 5916, 6101, 5804, 5707}

	aumenti := contaAumenti(rilevazioni)
	aumentiFin := contaAumentiFin(rilevazioni, m)

	fmt.Println(aumenti, aumentiFin)
}
