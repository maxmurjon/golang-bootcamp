package funcs

import "fmt"

func Unique(arr []int){

	dublicate:=make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		dublicate[arr[i]]=append(dublicate[arr[i]], i)

		if len(dublicate[arr[i]])>0{
			continue
		}else{
			for j := i; j < len(arr); j++ {
				if arr[i]==arr[j]{
					dublicate[arr[i]]=append(dublicate[arr[i]], j)
				}else{
					continue
				}
			}
		}
	}
	for key,v:=range dublicate{
		if len(v)==1{
			delete(dublicate,key)
		}else{
			continue
		}
	}
	fmt.Println(dublicate)
}