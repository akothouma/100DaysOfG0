package main

import("fmt")
 
func Max( s[]string)string{
	var integerSlice[]int
	
	for _,oneString:=range s{
		result:=0
		for _,oneRune:=range oneString{
			result=result*10 +int(oneRune)-'0'//atoi	
		}	
		integerSlice=append(integerSlice, result)		
	}
	for i:=0;i<len(integerSlice)-1;i++{ // swap to get max at last index
		if integerSlice[i+1]<integerSlice[i]{
			integerSlice[i+1],integerSlice[i]=integerSlice[i],integerSlice[i+1]
		}
	}
	 wantedInt:=integerSlice[len(integerSlice)-1]
	 var resultString, finalAnswer string 
	 for wantedInt>0{		
		var remainder int
		remainder=wantedInt%10
		wantedInt=wantedInt/10
        resultString+=string(rune(remainder)+'0') //itoa
	   }
	   for i:=0;i<len(resultString);i++{
		finalAnswer+=string(resultString[len(resultString)-i-1])//reversing string
	   }
	   return finalAnswer
 }

 func main(){
   str:=[]string{"520","1","1020","730","18","2"}
   finalAnswer:=Max(str)
   fmt.Println(finalAnswer)

 }