package funcs

import "fmt"

func Unique(arr []int) {

	dublicate := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		dublicate[arr[i]] = append(dublicate[arr[i]], i)
		for j := i; j < len(arr); j++ {
			if len(dublicate[arr[i]]) > 0 {
				continue
			}
			if arr[i] == arr[j] {
				dublicate[arr[i]] = append(dublicate[arr[i]], j)
			} else {
				continue
			}
		}

	}
	for key, v := range dublicate {
		if len(v) == 1 {
			delete(dublicate, key)
		} else {
			continue
		}
	}
	fmt.Println(dublicate)
}

func Shuffle(arr []int){
i1:=0
i2:=1
i3:=2
for true{
	fmt.Printf("%v %v %v\n",arr[i1],arr[i2],arr[i3])
		i1+=1
		i2+=1
		i3+=1
		if len(arr)==i1{
			i1=0
		}
		if len(arr)==i2{
			i2=0
		}
		if len(arr)==i3{
			i3=0
		}
}

}

func LenArray(arr []string)map[int]int{
	newMap:=make(map[int]int)
	for i,v:=range arr{
		space:=0
		for _,letter:=range v{
			if string(letter)==" "{
				space+=1
			}
		}
		newMap[i]=space
		space=0
	}
	return newMap
}