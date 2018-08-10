package main 

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	info := make(map[string][]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if !in(info[line], filename) {
				info[line] = append(info[line], filename)
			}
		}
	}
	for k, v := range info {
		if (len(v) > 1) {
			fmt.Printf("%s the dup file names are : ", k)
			for _, val := range v {
				fmt.Printf("%s ", val)
			}
			fmt.Printf("\n")
		}
	}

}


func in(data []string,tar string) bool {
	for _, val := range data {
		if tar == val {
			return true
		}
	}
	return false
}