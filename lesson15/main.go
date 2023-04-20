package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(array([]int{1,23,64,86,90}))
	fmt.Println(checkIfPangram("thequickbrowydog"))

}

func array(arr []int) (arr2 []int) {
	for i := 0; i < len(arr); i++ {
		if a := strconv.Itoa(arr[i]); len(a) > 1 {
			arr2 = append(arr2, (arr[i]+(i+1))%10)
		} else {
			arr2 = append(arr2, arr[i]+(i+1))
		}
	}

	return
}

func checkIfPangram(sentence string) bool {
	var alfabet [26]bool
	for _, v := range sentence {
		if v >= 'a' && v <= 'z' {
			alfabet[v-'a'] = true
		}
	}
	// fmt. Println(alfabet)
	for _, v := range alfabet {
		if !v {
			return false
		}
	}
	return true
}
