package main

import("fmt")
func main(){
	str1:="booa"
	str2:="fooaz"
     var newstring string
	for _,ch:=range str1{ //loop looks for non-unique characters
		for _,char:= range str2{
			if char==ch{
               newstring+=string(ch)
			   break
			}
		}
	}
	fmt.Println(newstring) //ooa

	str3:=str1+str2 
	fmt.Println(str3)
	var sliceWanted []rune

	for _,ch:=range str3{ 
		duplicate:=false
		for _,alreadyIn:= range sliceWanted{
			for _,same:=range newstring{ // checks to remove non-unique string
				if ch==alreadyIn || ch==same{
					duplicate=true
					break
				   }
			}
          
		}
		if !duplicate{
			sliceWanted=append(len(sliceWanted,ch))
		}
	
	}
	fmt.Println(sliceWanted)
}