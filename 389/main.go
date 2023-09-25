package main

import "fmt"


func main() {
	fmt.Println(findTheDifference("abc","abcd"))
}

func findTheDifference(s string, t string) byte {
	sumS:=0
	sumT:=0
	for _, v := range s {
		sumS+=int(v)
	}
	for _, v := range t {
		sumT+=int(v)
	}
	return byte(sumT-sumS)
}