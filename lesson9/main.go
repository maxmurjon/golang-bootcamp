package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var morseAlfa =make(map[string]string)
	sozInput,_:=bufio.NewReader(os.Stdin).ReadString('\n')
	sozInput=strings.ToUpper(sozInput)
	morseAlfa["A"]=".-"
	morseAlfa["B"]="-..."
	morseAlfa["C"]="-.-."
	morseAlfa["D"]="-.."
	morseAlfa["E"]="."
	morseAlfa["F"]="..-."
	morseAlfa["G"]="--."
	morseAlfa["H"]="...."
	morseAlfa["I"]=".."
	morseAlfa["J"]=".---"
	morseAlfa["K"]="-.-"
	morseAlfa["L"]=".-.."
	morseAlfa["M"]="--"
	morseAlfa["N"]="-."
	morseAlfa["O"]="---"
	morseAlfa["P"]=".--."
	morseAlfa["Q"]="--.-"
	morseAlfa["R"]=".-."
	morseAlfa["S"]="..."
	morseAlfa["T"]="-"
	morseAlfa["U"]="..-"
	morseAlfa["V"]="...-"
	morseAlfa["W"]=".--"
	morseAlfa["X"]="-..-"
	morseAlfa["Y"]="-.--"
	morseAlfa["Z"]="--.."
	morseAlfa["1"]=".----"
	morseAlfa["2"]="..---"
	morseAlfa["3"]="...--"
	morseAlfa["4"]="....-"
	morseAlfa["5"]="....."
	morseAlfa["6"]="-...."
	morseAlfa["7"]="--..."
	morseAlfa["8"]="---.."
	morseAlfa["9"]="----."
	morseAlfa["0"]="-----"

	soz:=""
	for _,v:=range sozInput{
		if morseAlfa[string(v)]!=""&&morseAlfa[string(v)]!=" "{
			soz+=morseAlfa[string(v)]
		}else if string(v)==" "{
			soz+=" "
		}
	}
	fmt.Println(soz)
}