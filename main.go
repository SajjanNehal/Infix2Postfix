package main

import(
	"fmt"
	"os"
	"strings"
)

var opstack Stack

func IsOperator(C uint8) bool{
	return strings.ContainsAny(string(C), "+ & - & * & /")
}

func IteratesString(input_infix string){
	for i := 0; i < len(input_infix); i++ {
		if IsOperator(input_infix[i]){
			if opstack.Empty(){
				opstack.Push(input_infix[i])
			}
			// fmt.Println(string(input_infix[i]))
		}else{
			fmt.Println(string(input_infix[i]))
		}
	}
}

func main(){
	user_input := os.Args[1]

	IteratesString(user_input)
}
