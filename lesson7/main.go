package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Golang-bootcamp/lesson7/funcs"
)

func main() {

	// funcs.Unique([]int{4,50,4,20,4,20})
	// funcs.Shuffle([]int{1,2,3,4,5,6,7})

	count := 0
	fmt.Printf("Nechta son kiritmoqchisiz ")
	fmt.Scanln(&count)
	slice1 := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("SO'z kiriting ")
		in,_:= bufio.NewReader(os.Stdin).ReadString('\n')
		slice1[i]=in
		fmt.Println(slice1)
	}
	fmt.Println(funcs.LenArray(slice1))
	
}
