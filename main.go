//WAP that takes a string and returns the wanted substring 

package main

import (
	"fmt"
)

func main() {
	strGiven := "erui,lkjhgfdsdsertdefiuytr"
	strWanted := "fds"
	sliceWanted := []rune{}
	var newstring,result string

	for _, ch := range strWanted {  //finding values that are wanted
		for _, char := range strGiven {
			if char == ch {
				newstring += string(ch)
			}
		}
	}
	
	for _, wanted:= range newstring { //removing repitition in wanted values
		duplicate := false
		for _, alreadyin := range sliceWanted {
			if wanted == alreadyin {
				duplicate = true
				break
			}
		}
		if !duplicate {
			sliceWanted = append(sliceWanted, wanted)
		}
	}

	for _,ch:= range sliceWanted{
        result+=string(ch) 
	}
	fmt.Println(result)

}
