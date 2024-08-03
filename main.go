package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not                  "))
	fmt.Print(LastWord("lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	if len(s) == 0 {
		return "\n"
	}
	word := ""
	reqS := []string{}
	for i, ch := range s {
		if ch == ' ' && len(word) == 0 {
			continue
		}
		if ch != ' ' {
			word += string(ch)
		}
		if ch == ' ' && i != len(s) {
			reqS = append(reqS, word)
			word = ""
		}
		if i == len(s)-1 {
			break
		}
	}

	if len(word) != 0 {
		reqS = append(reqS, word)
	}
	if len(reqS) == 0 {
		return "\n"
	}

	return reqS[len(reqS)-1] + "\n"
}
