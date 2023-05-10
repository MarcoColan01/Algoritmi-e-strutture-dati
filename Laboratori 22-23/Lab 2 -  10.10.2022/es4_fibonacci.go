package main

import (
	"fmt"
	"os"
	"strconv"
)

/*func f_rec (n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return f_rec (n -1) + f_rec (n -2)
}*/

func f_iter2 (n int) uint64 {
    var f uint64
    var f1 , f2 uint64 = 1, 1
    if n == 2 || n == 1 {
        return 1
}

    for i := 2; i <= n; i++ {
        n--
        f = f1 + f2
        f1 = f2
        f2 = f
    }
    return f
}

/*func f_iter1 (n int) uint64 {
    var f uint64
    var f1 , f2 uint64 = 1, 1
    if n == 2 || n == 1 {
        return 1
}

    for n >= 3 {
        n --
        f = f1 + f2
        f1 = f2
        f2 = f
    }
    return f
}*/

func main(){
	n,_:= strconv.Atoi(os.Args[1]);

	//fmt.Println("Il numero di fibonacci di",n,"è:", f_iter1(n));
	//fmt.Println("Il numero di fibonacci di",n,"è:", f_rec(n));
	fmt.Println("Il numero di fibonacci di",n,"è:", f_iter2(n));
}