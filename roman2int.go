package main

import (
	. "fmt"
)

func RomanChar2Int(ch uint8) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func RomanString2Int(roman string) int {
	sum := 0
	for i := len(roman) - 1; i >= 0; i-- {
		cur := roman[i]
		var prev uint8
		if i != 0 {
			prev = roman[i-1]
		} else {
			prev = cur
		}
		if RomanChar2Int(prev) < RomanChar2Int(cur) {
			sum += RomanChar2Int(cur) - RomanChar2Int(prev)
			i--
		} else {
			sum += RomanChar2Int(cur)
		}
	}
	return sum
}

func main() {
	Println(RomanString2Int("MCMXC"))
	Println(RomanString2Int("MMVIII"))
	Println(RomanString2Int("XCIX"))
	Println(RomanString2Int("XLVII"))
	Println(RomanString2Int("MMMMCMXCIX"))
}
