package main
import (
	"fmt"
)
func int_to_roman(n int) string {
	num := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	roman :=[] string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	result := ""

	for i:=0; i< len(num); i++ {
		for num[i]<=n {
			result += roman[i]
			n -= num[i]

		}
	}
	return result
}

func main() {
	n := 36
	r := int_to_roman(n)
	fmt.Printf("%d in Roman letter is %s", n, r)
}