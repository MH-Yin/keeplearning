package golang

import (
	"fmt"
)

// "word zz mtv"
func wordZz(input string) {
	var result string
	inputs := []rune(input)
	lens := len(input)
	var cache []rune
	for i := lens - 1; i >= 0; i-- {
		if inputs[i] == 32 {
			result = string(cache) + " " + result
			cache = []rune{}
			continue
		}
		cache = append(cache, inputs[i])
	}
	if len(cache) != 0 {
		result = string(cache) + " " + result
	}
	fmt.Println(result)
}
