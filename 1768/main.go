package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abcd","pq"))

}

func mergeAlternately(word1 string, word2 string) string {
    var text strings.Builder
	if len(word1)<=len(word2) {
		for i := 0; i < len(word1); i++ {
			text.WriteString(string(word1[i])+string(word2[i]))
		}
		if len(word1)<len(word2) {
			text.WriteString(word2[len(word1):])
		}
	}else{
		for i := 0; i < len(word2); i++ {
			text.WriteString(string(word1[i])+string(word2[i]))
		}
		if len(word2)<len(word1) {
			text.WriteString(word1[len(word2):])
		}
	}
	return text.String()
}