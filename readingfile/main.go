package main

import (
	"fmt"
	"strings"
	"log"
	"os"
)

func main() {
	inputContent, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	inputString:=string(inputContent)

	correctContent, err := os.ReadFile("correct.txt")
	if err != nil {
		log.Fatal(err)
	}
	correctString := string(correctContent)
	
	 if strings.Contains(correctString,inputString){
        os.WriteFile("result.txt",[]byte("Well Formatted\n"),0644)
	 }else{
		os.WriteFile("result.txt",correctContent,0644)
	 }
	 fmt.Println("result.txt has been generated")

}



