package main

import "fmt"

type node struct {
	sx, dx, up *node
	etichetta  rune
}

type tree *node

func findIndex(s string) int {
	count := 0
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] == '[' {
			count++
		} else if s[i] == ']' {
			count--
		} else if s[i] == ',' && count == 0 {
			break
		}
	}
	return i
}

func stringToTree(s string) tree {
	if len(s) == 0 {
		return nil
	}
	if s[0] == '[' {
		sx := stringToTree(s[1:])
		dx := stringHeigth()
		return &node{sx, dx, nil, rune(s[0])}
	} else if s[0] == ']' {
		return nil
	} else {
		return &node{nil, nil, nil, rune(s[0])}
	}
}

/*func treeToString(t tree) string {

}

func treeHight(t tree) int {

}*/

func stringHeigth(s string) int {
	cont := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			if cont > max {
				max = cont
			}
			cont = 0
		} else if s[i] == '[' {
			cont++
		}
	}
	return max
}

func dfs(t *node) {
	if t == nil {
		return
	}
	dfs(t.sx)
	fmt.Println(t.etichetta)
	dfs(t.dx)
}

func main() {
	s := "[[[[1,2],[3,4]],[5,6],[7,8]],9]"
	n := stringHeigth(s)
	fmt.Println(n)

	t := stringToTree(s)
	dfs(t)
}
