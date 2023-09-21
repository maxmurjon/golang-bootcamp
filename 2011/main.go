package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X","X++","X++"}))
}

func finalValueAfterOperations(operations []string) int {
	x:=0
	for i := range operations {
		if strings.Contains(operations[i],"++"){
			x++
		}else if strings.Contains(operations[i],"--"){
			x--
		}else{
			continue
		}
	}
	return x
}
