package harmony

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 Number Search
 Have the function NumberSearch(str) take the str parameter, search for all the numbers in the string, add them together,
 then return that final number divided by the total amount of letters in the string.
 For example: if str is "Hello6 9World 2, Nic8e D7ay!" the output should be 2.
 First if you add up all the numbers, 6 + 9 + 2 + 8 + 7 you get 32. Then there are 17 letters in the string. 32 / 17 = 1.882,
 and the final answer should be rounded to the nearest whole number, so the answer is 2.
 Only single digit numbers separated by spaces will be used throughout the whole string (So this won't ever be the case: hello44444 world).
 Each string will also have at least one letter.
 Examples
 Input: "H3ello9-9"
 Output: 4
 Input: "One Number*1*"
 Output: 0
*/
func TestNumberSearch(t *testing.T) {
	assert.Equal(t, 4, NumberSearch("H3ello9-9"))
	assert.Equal(t, 0, NumberSearch("One Number*1*"))
}

/*
 Distinct List
 Have the function DistinctList(arr) take the array of numbers stored in arr and determine the total number of duplicate entries.
 For example if the input is [1, 2, 2, 2, 3] then your program should output 2 because there are two duplicates of one of the elements.
 Examples
 Input: []int {0,-2,-2,5,5,5}
 Output: 3
 Input: []int {100,2,101,4}
 Output: 0
*/
func TestDistinctList(t *testing.T) {
	assert.Equal(t, 3, DistinctList([]int{0, -2, -2, 5, 5, 5}))
	assert.Equal(t, 0, DistinctList([]int{100, 2, 101, 4}))
}

/*
 Matrix Spiral
 Have the function MatrixSpiral(strArr) read the array of strings stored in strArr which will represent a 2D N matrix,
 and your program should return the elements after printing them in a clockwise, spiral order.
 You should return the newly formed list of elements as a string with the numbers separated by commas.
 For example: if strArr is "[1, 2, 3]", "[4, 5, 6]", "[7, 8, 9]" then this looks like the following 2D matrix:
  1 2 3
  4 5 6
  7 8 9
 So your program should return the elements of this matrix in a clockwise, spiral order which is: 1,2,3,6,9,8,7,4,5
 Examples
 Input: []string {"[1, 2]", "[10, 14]"}
 Output: 1,2,14,10
 Input: []string {"[4, 5, 6, 5]", "[1, 1, 2, 2]", "[5, 4, 2, 9]"}
 Output: 4,5,6,5,2,9,2,4,5,1,1,2
*/
func TestMatrixSpiral(t *testing.T) {
	assert.Equal(t, "1,2,14,10", MatrixSpiral([]string{"[1, 2]", "[10, 14]"}))
	assert.Equal(t, "4,5,6,5,2,9,2,4,5,1,1,2", MatrixSpiral([]string{"[4, 5, 6, 5]", "[1, 1, 2, 2]", "[5, 4, 2, 9]"}))
}
