package main

import "fmt"

//function returns the mirror of a letter eg a return z and z returns a ,s returns h and vice versa
func alphamirror( s string)string{
  stringToRange:=[]rune(s)

  for i,char:=range stringToRange{
	if char>='a'&& char<='z'{ //checks if character is small letter
		stringToRange[i]= 'z'-(char-'a')
	}else if char>='A'&& char<='Z'{ //checks if character is a capital letter
		stringToRange[i]='Z'-(char-'A')
	}else{ // character that is neither small nor capital 
		stringToRange[i]=char
	}
  }
  result:=string(stringToRange)
  return result
}

func main(){
	str:="Svool gl gsv hgzgfv Gln Nylbz."
	result:=alphamirror(str)
	fmt.Println(result)
}