package main

import(
	"fmt"
	"os"
	"strings"
)

var opstack Stack
var postfix string

func PopStackUntilEmpty(){
	for !opstack.Empty(){
		postfix += opstack.Top().(string)
		opstack.Pop()
	}
}

func GetOperatorWeight(op string) int{
	weight := -1
	switch op {
		case "+","-":
			weight = 1
			break
		case "*", "/":
			weight = 2
			break
		case "^":
			weight = 3
			break
	}
	return weight
}

func IsPresHigh(nextchar string, stacktop string) bool{
	nextchar_weight := GetOperatorWeight(nextchar)
	stacktop_weight := GetOperatorWeight(stacktop)

	if nextchar_weight == stacktop_weight{
		return true
	}else if nextchar_weight > stacktop_weight{
		return true
	}else{
		return false
	}
}

func PopStackUntilLeftParenthesis(){
	for opstack.Top().(string) != "(" {
		postfix += opstack.Top().(string)
		opstack.Pop()
	}
	opstack.Pop()
}

func IsOperator(C string) bool{
	return strings.ContainsAny(C, "+ & - & * & /")
}

func IteratesString(input_infix string){
	for i := 0; i < len(input_infix); i++ {
		if IsOperator(string(input_infix[i])){
			
			for !opstack.Empty() && opstack.Top().(string) != "(" && IsPresHigh(opstack.Top().(string), string(input_infix[i])){
				postfix += opstack.Top().(string)
				opstack.Pop()
			}
			opstack.Push(string(input_infix[i]))
		}else if string(input_infix[i]) == "(" {
			opstack.Push(string(input_infix[i]))
		}else if string(input_infix[i]) == ")" {
			PopStackUntilLeftParenthesis()
		}else{
			postfix += string(input_infix[i])
		}
	}
	PopStackUntilEmpty()
}

func main(){
	user_input := os.Args[1]

	IteratesString(user_input)

	fmt.Println(postfix)
}
