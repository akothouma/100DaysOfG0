package main

import "fmt"

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	var str string
	var duplicate bool
	count:=1
	for _, ch := range s {
		for _, ch1 := range str {
			duplicate=false
			if ch == ch1 && ch != ' ' {
				duplicate = true
				continue

			}
		}
		if duplicate {
		 count++
		}else{
			str +=Itoa(count)+string(ch)	
			count=1	
		}
		
	}

	return str
}

func Itoa(a int) string {
	return string(rune(a) + '0')
}
