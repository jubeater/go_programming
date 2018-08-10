package main 
import (
	"fmt"
	"os"
)
func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("the idx is %d, the arg is %s \n", idx, arg)
	}
}