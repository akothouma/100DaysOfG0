package main

import (
	"fmt"
)

func main(){
fmt.Println(UniqueChar("foo","boo"))
}

func UniqueChar(s1,s2 string)int{
	uniqueChars:=make(map[rune]int)

	for _,char:=range s1{
		uniqueChars[char]++
	}
	for _,char:=range s2{
		uniqueChars[char]++
	}
	count:=0
    for _,value:= range uniqueChars{
       if  value==1{
          count++
	   }
	}
	return count
}