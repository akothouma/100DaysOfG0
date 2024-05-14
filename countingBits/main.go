// WAP that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number.
// You can guarantee that input is non-negative.
// Example: The binary representation of 1234 is 10011010010, so the function should return 5
package main

func CountBits(r uint) int {
	binaryEqu := ""
	if r < 0 {
		return 0
	} else if r == 0 {
		return 0
	} else {
		var remainder uint
		for r > 0 {
			remainder = r % 2
			r = r / 2
			binaryEqu += string(rune(remainder + '0'))
		}
	}
	count := 0
	for _, ch := range binaryEqu {
		if ch == '1' {
			count++
		}
	}
	return count
}

func main() {
	CountBits(1234)
}
