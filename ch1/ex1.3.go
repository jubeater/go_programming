package main 

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {
	input := os.Args[1:]
	start := time.Now()
	prev(input)
	fmt.Printf("the prev solution used %.10f secs \n", time.Since(start).Seconds())
	second := time.Now()
	cur(input)
	fmt.Printf("the cur solution used %.10f secs \n", time.Since(second).Seconds())
}


func prev(input []string){
	var s string
	var sep string
	for _, arg := range input {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
}


func cur(input []string){
	fmt.Println(strings.Join(input, " "))
}