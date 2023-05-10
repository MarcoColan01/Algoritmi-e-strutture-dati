package main

import(
	"fmt"
	"strconv"
	"os"
) 

func multiply (x, y int) int {
	if y==1 {
		return x
	}else{
		return x + multiply (x, y -1)
	}
}

func main(){
	x,_:= strconv.Atoi(os.Args[1]);
	y,_:= strconv.Atoi(os.Args[2]);

	fmt.Println(x,"*",y,"=", multiply(x,y));
}