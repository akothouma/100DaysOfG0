//WAP to find repeated runes in a string and count number of repeat times 

package main
import "fmt"
func main(){
str:="asdfghjswtrbyuioskkjhgdetuiuwt"
 mapCount:=make(map[rune]int)
 

 for _,ch:=range str{
	mapCount[ch]++
 }
 for key,value:=range mapCount{
	if value>1{
		fmt.Printf("%s repeated %d times \n",string(key),value)
	 }
	 
 }

}