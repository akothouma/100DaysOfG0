package main

import "fmt"

func main(){
	fmt.Println(Rot13("hello there"))
}
func Rot13(s string) string {
	var char1 rune
	var answer string
	for _, char := range s {
		if char >= 97 && char <= 122 || char >= 65 && char <= 90 {
			if char < 109 || char < 77 {
				char1 = char + 13
			} else if char > 109 || char > 77 {
				char1 = char - 13
			} else if char == 109|| char == 77 {
				char1 = char -13
			}
		} else {
			char1 = char
		}
		answer += string(char1)
	}
	return answer
}
