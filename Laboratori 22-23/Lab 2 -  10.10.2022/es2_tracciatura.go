package main;

import "fmt";

func max(n int, m int) int{
	if n > m{
		return n
	}else{
		return m
	}
}

func largest (numbers []int) int {
	n := len(numbers)
	if n==1 {
		return numbers [0]
	}
	return max (numbers[n-1], largest(numbers[0:n-1]))
}

func main(){
	vet := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8, 110}
	fmt.Println("Il max nel vettore è:", largest(vet));
}