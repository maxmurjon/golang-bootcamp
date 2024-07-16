package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3,2,3}, 6))
}

func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i,num := range  nums{
		value:=target-num
		if index,ok:=hash[value]; ok {
			return []int{i,index}
		}

		hash[num]=i
	}
	return nil
}