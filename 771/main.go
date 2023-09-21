package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	jewelsNum := 0
	for i:=0;i<len(stones);i++ {
		if strings.Contains(jewels,string(stones[i])) {
			jewelsNum++
		}
	}
	return jewelsNum
}
