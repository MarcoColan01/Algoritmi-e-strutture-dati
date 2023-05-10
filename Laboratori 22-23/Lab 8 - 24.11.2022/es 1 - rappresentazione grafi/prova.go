package main 
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	filename := os.Args[1]
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan(){
		k := strings.Split(scanner.Text(), " ")
		fmt.Println(k[0], k[1], k[2:])
		//addUser(g, k[0], k[1], k[2:])
	}
}