package main

import "fmt"


func main() {
	str1 := "ddf6vewg64f"
	str2 := "gtwthgdwthdwfteewhrtag6h4ffdhsd"
	fmt.Println(Union(str1, str2))
}

func Union(s, s1 string) string{
	s3:=s+s1
	mp1 := make(map[rune]int)
	mp2 := make(map[rune]int)

	for _, ch := range s {
		mp1[ch]++
	}

	for _, ch := range s1 {
		mp2[ch]++
	}

	var arr string

	for key := range mp1 {
		arr+=string(key)
	}

	for key1 := range mp2 {
		if _, ok := mp1[key1]; !ok {
			arr+=string(key1)
		}
	}
fmt.Println(arr ,"here")
str:=""

	for _,ch:=range s3{
		for _,ch1:=range arr{
			if ch==ch1{
				if !isInside(ch,str){
					str+=string(ch)

				}

			}
		}
	
	}
	return str
}

func isInside(ch rune,s string)bool{
	for _,k:=range s{
		if k==ch{
			return true
		}
	}
	return false
}