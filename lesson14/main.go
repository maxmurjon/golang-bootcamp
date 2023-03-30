package main

import "fmt"

func main() {
	// fmt.Println(commonFactors(12,6))

	// words:=[]string{"notapalindrome","racecar"}
	// fmt.Println(firstPalindrome(words))

	// fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(flipAndInvertImage([][]int{ {0,1,1},{0,0,1},{0,1,0} }))

}

func commonFactors(a int, b int) int {
	var count int
	if a > b {
		for i := 1; i <= b; i++ {
			if a%i == 0 && b%i == 0 {
				count += 1
			}
		}
	} else {
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				count += 1
			}
		}
	}
	return count
}

func firstPalindrome(words []string) string {
	for _, v := range words {
		if v == reverse(v) {
			return v
		} else {
			continue
		}
	}
	return ""
}

func reverse(str string) string {
	strRet := ""
	for i := len(str) - 1; i >= 0; i-- {
		strRet += string(str[i])
	}
	return strRet
}

func selfDividingNumbers(left int, right int) []int {
	arr := []int{}
	for i := left; i <= right; i++ {
		if isDividable(i) {
			arr = append(arr, i)
		}
	}
	return arr
}
func isDividable(i int) bool {
	num := i
	for num > 0 {
		digit := num % 10
		if digit == 0 || i%digit != 0 {
			return false
		}
		num /= 10

	}
	return true
}

func flipAndInvertImage(image [][]int)[][]int{
	newSlice:=[][]int{}
	for _,v:=range image{
		slice1:=[]int{}
		for i := len(v)-1; i >=0; i-- {
			if v[i]==0{
				slice1 = append(slice1, 1)
			}else if v[i]==1{
				slice1=append(slice1, 0)
			}
		}
		newSlice = append(newSlice, slice1)
	}
	return newSlice
}