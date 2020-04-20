package main

import (
	"fmt"
)

var test_array = [8][3]string{
	{"a+b*c+d-e+f","abc*+d+e-f+","true"},
	{"(a+b)*(c+d)","ab+cd+*","true"},
	{"a*b+c*d","ab*cd*+","true"},
	{"a+b+c+d","ab+c+d+","true"},
	{"3+4*5/6","345*6/+","true"},
	{"(300+23)*(43-21)/(84+7)","30023+4321-*847+/","true"},
	{"a*b-c", "abc-*","false"},
	{"3+4*5/6","3456*/+","false"},
}

var test_pass = true

func Compare(test string, output string) string{
	if test == output{
		return "true"
	}else{
		return "false"
	}
}

func RunTest(){
	for i := 0; i < len(test_array); i++ {
		program_output := IteratesString(test_array[i][0])

		if Compare(test_array[i][1], program_output) == test_array[i][2]{
			fmt.Println("[+] Following Test is Passing:")
			fmt.Println("[+] Test output: "+ program_output + " Macthed with: " + test_array[i][1] + ". Should be " + test_array[i][2])
		}else{
			test_pass = false
			fmt.Println("[!] Following Test is failing:")
			fmt.Println("[!] Test output: "+ program_output + " Macthed with: " + test_array[i][1] + ". Should be " + test_array[i][2])
		}
	}

}

func main(){
	fmt.Println("[*] Runing Tests")

	RunTest()

	if test_pass == true{
		fmt.Println("All Test Passed.")
	}else{
		fmt.Println("Test Failed.")
	}
}
