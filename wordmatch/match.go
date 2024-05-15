package main

import (
	"fmt"
	"os"
)


func main(){
	input1:=os.Args[1]
	input2:=os.Args[2]
	result:=""

	for _,char2:= range input2{
		for i,char1:= range input1{
			if char2==char1{
				result+=string(char1)
				input1=input1[i+1:]
				break

			}
		}
	}
	fmt.Println(result)
}