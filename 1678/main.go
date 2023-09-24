package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(interpretTwo("G()(al)"))
}

// Fist soluiton in this problem
func interpret(command string) string {
	text := command
	*&text=strings.ReplaceAll(text, "()", "o")
	*&text=strings.ReplaceAll(text, "(al)", "al")
	return text
}

// Second solution this problem
func interpretTwo(command string) string {
    var text strings.Builder
	i:=0
    for i<len(command){
        if command[i] == 'G'{
            text.WriteByte('G')
            i++
        }else if command[i:i+2]=="()"{
            text.WriteString("o")
            i+=2
        }else if command[i:i+4]=="(al)"{
            text.WriteString("al")
            i+=4
        }
    }
    return text.String()
}